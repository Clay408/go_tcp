package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgData() []byte
	GetMsgLen() uint32

	SetMsgId(uint32)
	SetMsgData([]byte)
	SetMsgLen(uint32)
}
