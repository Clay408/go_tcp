package znet

import (
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

	//3.阻塞等待客户端连接，处理客户端连接业务(读写)
	for {
		//等待客户端连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}

		//已经与客户端建立连接,做一个最大512字节长度的回显业务
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf) // 阻塞
				if err != nil {
					fmt.Println("receive buf err: ", err)
					continue
				}

				//回显功能
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write back err: ", err)
					continue
				}
			}
		}()
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
