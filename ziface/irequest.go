package ziface

type IRequest interface {
	// GetConnection 获取当前请求对应的客户端链接
	GetConnection() IConnection

	// GetData 获取请求数据
	GetData() []byte

	// GetDataLength 获取数据长度
	GetDataLength() uint32

	// GetMsgId 获取消息ID
	GetMsgId() uint32
}
