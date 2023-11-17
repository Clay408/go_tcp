package routers_test

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
	fmt.Println("==================== receive from client ", string(request.GetData()))
	response, err := request.GetConnection().GetDataPackHandle().Pack(znet.NewMessage(request.GetMsgId(), []byte("这里是PingHandle")))
	if err != nil {
		fmt.Println("Pack client ResponseMsg err ", err)
		return
	}
	_, err = request.GetConnection().GetTCPConnection().Write(response)
	if err != nil {
		fmt.Println("Response client err ", err)
		return
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("==================== receive from client ", string(request.GetData()))
	response, err := request.GetConnection().GetDataPackHandle().Pack(znet.NewMessage(request.GetMsgId(), []byte("这里是HelloHandler")))
	if err != nil {
		fmt.Println("Pack client ResponseMsg err ", err)
		return
	}
	_, err = request.GetConnection().GetTCPConnection().Write(response)
	if err != nil {
		fmt.Println("Response client err ", err)
		return
	}
}

func TestServer(t *testing.T) {
	server := znet.NewServer("123")
	server.AddRouter(uint32(0), &HelloRouter{})
	server.AddRouter(uint32(1), &PingRouter{})
	server.Serve()
}
