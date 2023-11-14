package znet

import (
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

// Start 启动链接(让当前的链接开始准备工作)
func (c *Connection) Start() {

}

// Stop 停止链接(结束当前连接的工作)
func (c *Connection) Stop() {

}

// GetTCPConnection 获取当前链接所绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return nil
}

// GetConnID 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return 1
}

// RemoteAddr 获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return nil
}

// Send 发送数据，将数据发送给远程的客户端
func (c *Connection) Send([]byte) error {
	return nil
}
