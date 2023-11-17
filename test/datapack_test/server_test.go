package datapack_test

import (
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"github.com/Clay408/zinx/znet"
	"testing"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Client send msg ", string(request.GetData()))
	err := request.GetConnection().Send(1, []byte("ping....ping....ping...."))
	if err != nil {
		fmt.Println("Write to client err ", err)
		return
	}
}

func TestStartRouterServer(t *testing.T) {
	server := znet.NewServer("PackUnPackServer")
	server.AddRouter(uint32(1), &PingRouter{})
	server.Serve()
}
