package ziface

type IMsgHandle interface {
	// DoMsgHandle 业务处理
	DoMsgHandle(request IRequest)
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)

	// GetCurrentRouterNum 获取当前服务器的路由数量
	GetCurrentRouterNum() int
}
