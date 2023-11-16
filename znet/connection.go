package znet

import (
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"net"
)

// Connection 链接模块
type Connection struct {
	Conn      *net.TCPConn      //当前链接的套接字
	ConnId    uint32            //当前链接ID
	IsClosed  bool              //当前链接是否关闭
	HandleAPI ziface.HandleFunc // 当前链接的业务处理
	ExitChan  chan bool         //告知当前链接退出或者停止的channel
}

func NewConnection(conn *net.TCPConn, connId uint32, callBackAPI ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connId,
		HandleAPI: callBackAPI,
		IsClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// StartReader 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running .....")
	defer fmt.Printf("ConnID=%d, Reader is exit ,remote addr is %s", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buff中,目前最大就是512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}

		//调用当前链接所绑定的HandlerAPI
		if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnectID", c.ConnId, "Handler err", err)
			break
		}
	}
}

// Start 启动链接(让当前的链接开始准备工作)
func (c *Connection) Start() {
	fmt.Println("Start connection ,ConnID=", c.ConnId)
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

// Send 发送数据，将数据发送给远程的客户端
func (c *Connection) Send([]byte) error {
	return nil
}
