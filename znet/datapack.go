package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/Clay408/zinx/utils"
	"github.com/Clay408/zinx/ziface"
)

type DataPack struct {
}

func (d *DataPack) GetHeadLen() uint32 {
	// 4字节数据长度 | 4字节消息ID
	return 8
}

// Pack 将数据打包成字节流
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放字节数据的缓冲区
	dataBuf := bytes.NewBuffer([]byte{})
	//将dataLen写入dataBuf中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将msgId写入
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//写入数据
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuf.Bytes(), nil
}

// UnPack 拆包(将包的head信息读取出来)
func (d *DataPack) UnPack(headInfo []byte) (ziface.IMessage, error) {
	//创建一个输入二进制数据的IOReader
	dataBuf := bytes.NewReader(headInfo)
	//只解压head信息，得到dataLen和msgId
	msg := &Message{}

	//先读取长度
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.dataLen); err != nil {
		return nil, err
	}
	//读取消息ID
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}

	//判断dataLen是否已经超过了我们限定的最大数据长度
	if utils.ServerConfig.MaxPackages > 0 && msg.dataLen > utils.ServerConfig.MaxPackages {
		return nil, errors.New("too Large data receive")
	}

	//读取数据
	return msg, nil
}
