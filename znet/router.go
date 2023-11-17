package znet

import (
	"github.com/Clay408/zinx/ziface"
)

// BaseRouter 路由的基类，默认空实现路由方法，实现类可以有选择的实现其中的方法
type BaseRouter struct {
}

// PreHandle 业务前置处理
func (b *BaseRouter) PreHandle(request ziface.IRequest) {

}

// Handle 业务处理
func (b *BaseRouter) Handle(request ziface.IRequest) {
}

// PostHandle 业务后置处理
func (b *BaseRouter) PostHandle(request ziface.IRequest) {
}
