package znet

import "github.com/lorenzoyu2000/zinx/ziface"

type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request ziface.IRequest)  {}
func (b *BaseRouter) Handle(request ziface.IRequest)     {}
func (b *BaseRouter) PostHandle(request ziface.IRequest) {}
