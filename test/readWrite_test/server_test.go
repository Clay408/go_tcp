package readWrite_test

import (
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"github.com/Clay408/zinx/znet"
	"testing"
)

type BisRouter struct {
	znet.BaseRouter
}

// 重写业务处理逻辑
func (b *BisRouter) Handle(request ziface.IRequest) {
	conn := request.GetConnection()
	msgId := request.GetMsgId()
	fmt.Println("收到客户端发来的数据")
	//客户端响应
	conn.Send(msgId+1, []byte("BisRouter handler received msg "))
}

func TestStartServer(t *testing.T) {
	server := znet.NewServer("")
	server.AddRouter(uint32(1), &BisRouter{})
	server.Serve()
}
