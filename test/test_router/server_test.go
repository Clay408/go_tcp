package test_router

import (
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"github.com/Clay408/zinx/znet"
	"testing"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	//给客户端写回数据
	request.GetConnection().GetTCPConnection().Write([]byte("before ping .... \n"))
	fmt.Println("PreHandle Receive from client:  ", request.GetData())
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("ping .... \n"))
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("after ping .... \n"))
}

func TestStartRouterServer(t *testing.T) {
	server := znet.NewServer("router server")
	server.AddRouter(&PingRouter{})
	server.Serve()
}
