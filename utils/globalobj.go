package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Clay408/zinx/ziface"
	"io/ioutil"
)

// 全局配置参数
type serverConfig struct {

	/**
	Server 相关配置
	*/
	TcpServer ziface.IServer //当前zinx服务的全局的server对象
	Host      string         //主机
	TcpPort   int            //端口
	Name      string         //服务器名称

	/**
	Zinx 相关配置
	*/

	Version     string //zinx版本号
	MaxConn     int    //最大链接数
	MaxPackages uint32 // 最大数据包长度，超过这个长度就会进行拆包
}

// 定义一个全局的对外对象
var ServerConfig *serverConfig

// init 初始化加载全局配置
func init() {
	//如果配置文件没有加载，默认的值
	ServerConfig = &serverConfig{
		Name:        "ZinxServerApp",
		Version:     "V1.0",
		TcpPort:     8999,
		MaxConn:     1000,
		MaxPackages: 4096,
	}
	//尝试从配置文件中加载这些配置信息
	ServerConfig.LoadConfig()
}

func (config *serverConfig) LoadConfig() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read config file err, use default config")
	}
	//将json文件数据解析到 serverConfig中
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
