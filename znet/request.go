package znet

import "github.com/Clay408/zinx/ziface"

type Request struct {
	conn   ziface.IConnection //客户端请求的链接
	data   []byte             //客户端请求的数据
	length int                //客户端请求数据长度
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
