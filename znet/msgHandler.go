package znet

import (
	"fmt"
	"github.com/Clay408/zinx/utils"
	"github.com/Clay408/zinx/ziface"
)

type MsgHandle struct {
	Apis           map[uint32]ziface.IRouter
	WorkerPoolSize uint32                 //业务工作Worker池的数量
	TaskQueue      []chan ziface.IRequest //Worker负责取任务的消息队列
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.ServerConfig.WorkerPoolSize,
		//一个worker对应一个queue
		TaskQueue: make([]chan ziface.IRequest, utils.ServerConfig.WorkerPoolSize),
	}
}

// DoMsgHandle 业务处理
func (m *MsgHandle) DoMsgHandle(request ziface.IRequest) {
	msgId := request.GetMsgId()
	handler, ok := m.Apis[msgId]
	if !ok {
		fmt.Println("Router handle is not exist,You need add it.")
		request.GetConnection().Send(msgId, []byte("Router handle is not exist,You need add it."))
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

// StartWorkerPool 启动worker工作池
func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		//给当前worker对应的任务队列开辟空间
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.ServerConfig.MaxWorkerTaskLen)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

// SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//根据ConnId来决定交给那个Worker处理，轮询分配
	workerId := request.GetConnection().GetConnID() % m.WorkerPoolSize
	//将请求消息发送给任务队列
	m.TaskQueue[workerId] <- request
}

// StartOneWorker 启动一个工作线程和对应的工作队列
func (m *MsgHandle) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Printf("WorkerId:  %d is started \n", workerId)

	//不断地等待队列中的请求消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			m.DoMsgHandle(request)
		}
	}
}
