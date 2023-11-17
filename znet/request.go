package znet

import "github.com/Clay408/zinx/ziface"

type Request struct {
	conn ziface.IConnection //客户端请求的链接
	msg  ziface.IMessage    //消息
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetDataLength() uint32 {
	return r.msg.GetMsgLen()
}
