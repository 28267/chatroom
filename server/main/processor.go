package main //总控器

import (
	// "encoding/json"
	"fmt"
	"io"

	"github.com/28267/chatroom/common/message"
	process2 "github.com/28267/chatroom/server/process"
	"github.com/28267/chatroom/server/utils"

	// "io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (pro *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessLogin(mes)
		//处理登录
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessRegister(mes)
		//处理注册
	default:
		fmt.Println("该消息类型不存在，无法处理")
	}
	return
}

func (processor *Processor) Process2() (err error) {
	defer processor.Conn.Close()
	//循环的接收客户端发送的消息
	for {
		tf := &utils.Transfer{
			Conn: processor.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出..")
				return err
			} else {
				fmt.Println("tf.readPkg(conn)-错误", err)
				return err
			}
		}
		err = processor.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}
