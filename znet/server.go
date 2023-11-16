package znet

import (
	"errors"
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"net"
)

// Server 定义一个Server服务器模块
type Server struct {
	Name      string //服务器名称
	IPVersion string //服务器监听的IP版本
	Ip        string //服务器IP地址
	Port      int    //服务器监听的端口
}

// CallBack 定义当前客户端连接所绑定的API(目前是写成固定的，后续应该由Demo端自定义)
func CallBack(conn *net.TCPConn, data []byte, cnt int) error {
	//回显的业务
	fmt.Println("【Conn Handle call back to client】")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write data err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start Server Listener at IP %s ,Port: %d is starting\n]", s.Ip, s.Port)
	//1.获取一个TCP的Addr
	tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}
	//2.监听服务器地址
	listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
	if err != nil {
		fmt.Print("listen", s.IPVersion, ", error", err)
		return
	}

	fmt.Printf("start zinx server success, %s is listening...... \n", s.Name)

	//分配ConnID
	var cid uint32
	cid = 0

	//3.阻塞等待客户端连接，处理客户端连接业务(读写)
	for {
		//等待客户端连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}

		//将处理新连接的业务方法和Conn进行绑定得到Conn模块
		dealConn := NewConnection(conn, cid, CallBack)
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

	//阻塞
	select {}
}

// NewServer 初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      name,
		IPVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
	}
	return server
}
