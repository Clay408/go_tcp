package znet

type Message struct {
	id      uint32 //消息的ID
	dataLen uint32 //消息长度
	data    []byte //消息实体
}

func NewMessage(msgId uint32, data []byte) *Message {
	msg := &Message{
		id:      msgId,
		dataLen: uint32(len(data)),
		data:    data,
	}
	return msg
}

func (m *Message) GetMsgId() uint32 {
	return m.id
}
func (m *Message) GetMsgData() []byte {
	return m.data
}
func (m *Message) GetMsgLen() uint32 {
	return m.dataLen
}
func (m *Message) SetMsgId(id uint32) {
	m.id = id
}
func (m *Message) SetMsgData(data []byte) {
	m.data = data
}
func (m *Message) SetMsgLen(len uint32) {
	m.dataLen = len
}
