package ziface

/*
	消息管理抽象层
*/
type IMsgHandle interface {
	// 调度、执行消息对应的处理方法
	DoMsgRouter(IRequest)
	// 添加具体的处理逻辑
	AddRouter(uint32, IRouter)
	// 启动工作池
	StartWorkPool()
	// 发送消息到TaskQueue中
	SendMsgToTaskQueue(IRequest)
}
