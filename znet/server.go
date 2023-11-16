package znet

import (
	"fmt"
	"github.com/Clay408/zinx/utils"
	"github.com/Clay408/zinx/ziface"
	"net"
)

// Server 定义server服务器模块
type Server struct {
	Name      string         //服务器名称
	IPVersion string         //服务器监听的IP版本
	Ip        string         //服务器IP地址
	Port      int            //服务器监听的端口
	Router    ziface.IRouter //当前的Server添加一个router，server注册的链接对应的处理业务
	Exited    chan int       //异常退出标识通道
}

func (s *Server) Start() {
	//Router 必须要添加
	if s.Router == nil {
		fmt.Println("The router can not empty,Please add Router!!!!")
		s.Exited <- 1
		return
	}

	fmt.Printf("[Start Server Listener at IP %s ,Port: %d is starting\n]", s.Ip, s.Port)
	//获取一个TCP的Addr
	tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}
	//监听服务器地址
	listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
	if err != nil {
		fmt.Print("listen", s.IPVersion, ", error", err)
		return
	}

	fmt.Printf("start server success, %s is listening...... \n", s.Name)

	//分配ConnID
	var cid uint32
	cid = 0

	for {
		//阻塞等待客户端连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		fmt.Printf("client connected,connId: %s", string(cid))

		//将处理新连接的业务方法和Conn进行绑定得到Conn模块
		dealConn := NewConnection(conn, cid, s.Router)
		cid++
		//启动当前链接的业务处理
		go dealConn.Start()
	}
}

func (s *Server) Stop() {
	//TODO 将一些服务器的资源和状态，或者一些已经开辟的连接信息进行停止或者回收
}

func (s *Server) Serve() {
	go s.Start()

	//做一些启动服务器之后的额外工作.....
	//....

	//服务端等待指令，go程可以通过服务端的channel来与go服务端主程进行通信
	for {
		select {
		case <-s.Exited:
			fmt.Println("Server exited...")
			break
		}
	}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

// NewServer 初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      utils.ServerConfig.Name,
		IPVersion: "tcp4",
		Ip:        utils.ServerConfig.Host,
		Port:      utils.ServerConfig.TcpPort,
		Router:    nil,
		Exited:    make(chan int),
	}
	return server
}
