package ziface

type IDataPack interface {
	// GetHeadLen 获取包头的数据长度
	GetHeadLen() uint32
	// Pack 封包
	Pack(msg IMessage) ([]byte, error)
	// UnPack 拆包
	UnPack([]byte) (IMessage, error)
}
