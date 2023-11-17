package common

import (
	"github.com/Clay408/zinx/znet"
	"testing"
)

func TestServer(t *testing.T) {
	server := znet.NewServer("测试服务器")
	router := znet.BaseRouter{}
	server.AddRouter(uint32(1), &router)
	server.Serve()
}
