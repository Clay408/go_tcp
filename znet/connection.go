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
	Conn     *net.TCPConn     //当前链接的套接字
	ConnId   uint32           //当前链接ID
	IsClosed bool             //当前链接是否关闭
	ExitChan chan bool        //告知当前链接退出或者停止的channel
	Router   ziface.IRouter   //当前链接的业务处理路由
	packet   ziface.IDataPack //数据的处理方式
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connId,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	c.packet = &DataPack{}
	return c
}

// StartReader 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running .....")
	defer fmt.Printf("ConnID=%d, Reader is exit ,remote addr is %s", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取消息
		headData := make([]byte, c.packet.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			break
		}
		//拆包，得到msgId和msgLen
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

		//当前链接的业务处理(前 中 后)
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// Start 启动链接(让当前的链接开始准备工作)
func (c *Connection) Start() {
	fmt.Printf("start connection handle, connectionid: %s", string(c.ConnId))
	//启动从当前链接的读数据的业务
	go c.StartReader()

	//TODO 启动当前链接的写数据的业务
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

	//将数据发送给客户端
	if _, err := c.Conn.Write(binary); err != nil {
		fmt.Println("Write msg id ", msgId, "error: ", err)
		return errors.New("conn Write error")
	}
	return nil
}
