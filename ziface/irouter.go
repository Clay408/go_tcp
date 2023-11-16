package ziface

// 路由的抽象层
type IRouter interface {
	// PreHandle 业务处理之前
	PreHandle(request *IRequest)
	// Handle 业务处理
	Handle(request *IRequest)
	// PostHandle 业务处理之后
	PostHandle(request *IRequest)
}
