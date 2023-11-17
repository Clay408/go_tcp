package datapack

import (
	"fmt"
	"github.com/Clay408/zinx/znet"
	"io"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	//拨号
	conn, err := net.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println("connected to server err ", err)
		return
	}

	//开启服务端消息监听
	go StartReadData(conn)

	dp := &znet.DataPack{}

	for {
		time.Sleep(3 * time.Second)
		//发送封包的Message
		binary, err := dp.Pack(znet.NewMessage(0, []byte("Hello server!!!!")))
		if err != nil {
			fmt.Println("Pack data err ", err)
			break
		}
		//向服务端发送消息
		_, err = conn.Write(binary)
	}

}

// 监听服务端消息
func StartReadData(conn net.Conn) {
	fmt.Println("等待服务端消息....")
	dp := &znet.DataPack{}

	for {
		headInfo := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, headInfo)
		if err != nil {
			fmt.Println("Read headInfo err ", err)
			continue
		}

		msg, err := dp.UnPack(headInfo)
		if err != nil {
			fmt.Println("UnPack HeadInfo err ", err)
			continue
		}
		//根据长度读取真正的消息内容
		var data []byte
		data = make([]byte, msg.GetMsgLen())
		io.ReadFull(conn, data)

		fmt.Println("Receive from server msg ", string(data))
	}

}
