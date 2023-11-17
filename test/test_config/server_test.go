package test_router

import (
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"github.com/Clay408/zinx/znet"
	"testing"
)

type PingRouter1 struct {
	znet.BaseRouter
}

func (p *PingRouter1) PreHandle(request ziface.IRequest) {
	//给客户端写回数据
	request.GetConnection().GetTCPConnection().Write([]byte("before ping .... \n"))
	fmt.Printf("PreHandle Receive from client: %s , dataLength: %d", string(request.GetData()), request.GetDataLength())
}

func (p *PingRouter1) Handle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("ping .... \n"))
}

func (p *PingRouter1) PostHandle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("after ping .... \n"))
}

func TestStartRouterServer(t *testing.T) {
	server := znet.NewServer("router server")
	server.AddRouter(uint32(1), &PingRouter1{})
	server.Serve()
}
