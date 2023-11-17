package znet

import (
	"fmt"
	"github.com/Clay408/zinx/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	m := &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
	return m
}

// DoMsgHandle 业务处理
func (m *MsgHandle) DoMsgHandle(request ziface.IRequest) {
	msgId := request.GetMsgId()
	handler, ok := m.Apis[msgId]
	if !ok {
		dp := &DataPack{}
		fmt.Println("Router handle is not exist,You need add it.")
		errMsg, _ := dp.Pack(NewMessage(msgId, []byte("Router handle is not exist,You need add it.")))
		request.GetConnection().GetTCPConnection().Write(errMsg)
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 添加路由
func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		fmt.Println("MsgId", msgId, "Exist!")
		return
	}
	m.Apis[msgId] = router
	fmt.Println("Add Router success!!!")
}

func (m *MsgHandle) GetCurrentRouterNum() int {
	return len(m.Apis)
}
