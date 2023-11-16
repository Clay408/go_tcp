package ziface

type IRequest interface {
	// GetConnection 获取当前请求对应的客户端链接
	GetConnection() IConnection

	// GetData 获取请求数据
	GetData() []byte
}
