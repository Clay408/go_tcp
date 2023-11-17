package routers

import (
	"fmt"
	"github.com/Clay408/zinx/znet"
	"io"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMessage(0, []byte("msgId为1 的消息")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}

		//将headData字节流 拆包到msg中
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetMsgLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			data := make([]byte, msg.GetMsgLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			msg.SetMsgData(data)

			fmt.Println("==> Recv Msg: ID=", msg.GetMsgId(), ", len=", msg.GetMsgLen(), ", data=", string(msg.GetMsgData()))
		}

	}

}
