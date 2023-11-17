package readWrite

import (
	"fmt"
	"github.com/Clay408/zinx/znet"
	"math/rand"
	"net"
	"testing"
	"time"
)

func TestMultiClient(t *testing.T) {
	//拨号连接
	conn, err := net.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println("connect server error ", err)
		return
	}

	go readFromServer(conn)

	dp := znet.NewDataPack() //数据打包工具
	for {
		time.Sleep(1 * time.Second)
		//向服务端写数据
		msg := []byte("你好服务端")
		//msgId作为处理路由的映射关系，就像请求路径一样
		sendMsg := znet.NewMessage(1, msg)
		requestBinary, err := dp.Pack(sendMsg)
		if err != nil {
			fmt.Println("Pack data error ", err)
			continue
		}
		if _, err := conn.Write(requestBinary); err != nil {
			fmt.Println("Write to server error ", err)
			continue
		}

	}
}

func readFromServer(conn net.Conn) {
	for {
		response := make([]byte, 4096)
		cnt, err := conn.Read(response)
		if err != nil {
			fmt.Println("Read msg from server error ", err)
			continue
		}
		fmt.Println("Read server msg ", string(response[:cnt]))
	}

}

func RandomNum() int {
	// 设置种子，以确保每次运行生成的随机数都不同
	rand.NewSource(time.Now().UnixNano())
	// 生成一个10以内的随机数
	randomNumber := rand.Intn(10)
	return randomNumber
}
