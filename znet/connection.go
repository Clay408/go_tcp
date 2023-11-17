package znet

import (
	"errors"
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"io"
	"net"
)

// Connection 链接模块
type Connection struct {
	Conn     *net.TCPConn      //当前链接的套接字
	ConnId   uint32            //当前链接ID
	IsClosed bool              //当前链接是否关闭
	ExitChan chan bool         //告知当前链接退出或者停止的channel
	msgChan  chan []byte       //无缓冲管道，用于读、写两个goroutine之间的消息通信
	packet   ziface.IDataPack  //数据的处理方式
	handler  ziface.IMsgHandle //多路由处理器
}

func NewConnection(conn *net.TCPConn, connId uint32, handler ziface.IMsgHandle, dp ziface.IDataPack) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connId,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
		msgChan:  make(chan []byte),
		handler:  handler,
		packet:   dp,
	}
	return c
}

// Start 启动链接(让当前的链接开始准备工作)
func (c *Connection) Start() {
	fmt.Printf("start connection handle, connectionid: %s", string(c.ConnId))
	//启动从当前链接的读数据的业务
	go c.StartReader()
	//启动当前链接的写数据的业务
	go c.StartWriter()
}

// Stop 停止链接(结束当前连接的工作)
func (c *Connection) Stop() {
	fmt.Println("Connection stop .. ConnectionId = ", c.ConnId)

	//如果当前链接已经关闭
	if c.IsClosed == true {
		return
	}

	c.IsClosed = true
	//调用关闭Socket连接
	c.Conn.Close()
	//关闭管道
	close(c.ExitChan)
}

// GetTCPConnection 获取当前链接所绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetDataPackHandle() ziface.IDataPack {
	return c.packet
}

// GetConnID 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

// RemoteAddr 获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据，将数据发送给客户端
func (c *Connection) Send(msgId uint32, data []byte) error {
	if c.IsClosed == true {
		return errors.New("connection closed when send msg")
	}

	binary, err := c.packet.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("package msg error")
	}

	//将数据发送给客户端(触发Connection的Writer goroutine)
	c.msgChan <- binary
	return nil
}

// StartReader 接收客户端数据
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running .....")
	defer fmt.Printf("ConnID=%d, Reader is exit ,remote addr is %s", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()

	for {
		//按照协议格式读取数据头部内容
		headData := make([]byte, c.packet.GetHeadLen())
		//从指定连接通道中把指定的字节数组塞满(没有数据的情况下会阻塞在这里)
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			break
		}

		//对头部数据进行拆包，得到msgId和DataLen
		msg, err := c.packet.UnPack(headData)
		if err != nil {
			fmt.Println("UnPack err", err)
			break
		}

		//根据dataLen，再次读取Data， 放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		msg.SetMsgData(data)
		//构造消息请求对象
		req := Request{
			conn: c,
			msg:  msg,
		}

		//当前链接的业务处理
		go func(request ziface.IRequest) {
			c.handler.DoMsgHandle(request)
		}(&req)
	}

	c.ExitChan <- true
}

// StartWriter 向客户端返回响应数据
func (c *Connection) StartWriter() {
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")
	for {
		select {
		//等待的是包装好的数据，直接写回即可(触发点在Send方法中)
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
			fmt.Printf("%s[ConnID] Writer to client success", string(c.ConnId))
		case <-c.ExitChan:
			//conn已经关闭
			return
		}
	}
}
