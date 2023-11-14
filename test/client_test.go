package test

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	fmt.Println("client start")

	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client connect err: ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("你好服务器"))
		if err != nil {
			fmt.Println("send to server failed: ", err)
			return
		}

		//从服务器读取消息
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read from server err: ", err)
			return
		}

		fmt.Printf("server call back : %s , cnt=%d\n", buf, cnt)

		//阻塞
		time.Sleep(time.Second * 1)
	}

}
