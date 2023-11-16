package test_router

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClientSend(t *testing.T) {
	//拨号
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("connected to server err ", err)
		return
	}

	for {
		time.Sleep(1 * time.Second)
		_, err := conn.Write([]byte("Hello Server"))
		if err != nil {
			fmt.Println("Write to server err ", err)
			continue
		}

		//从服务器读取消息
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read from server err: ", err)
			return
		}
		fmt.Printf("receive from server : %s , cnt=%d\n", buf, cnt)
	}

}
