package ziface

type IMsgHandle interface {
	// DoMsgHandle 业务处理
	DoMsgHandle(request IRequest)
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)

	// GetCurrentRouterNum 获取当前服务器的处理器数量
	GetCurrentRouterNum() int
	// StartWorkerPool 启动工作线程池
	StartWorkerPool()
	// SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
	SendMsgToTaskQueue(request IRequest)
}
