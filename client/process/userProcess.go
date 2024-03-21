package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/28267/chatroom/client/utils"
	"github.com/28267/chatroom/common/message"
)

type UserProcess struct {
	//暂不需要字段
}

func (userProcess *UserProcess) Login(userId int, userPwd string) (err error) {

	// fmt.Printf("userId=%v,userPwd=%v",userId,userPwd)
	// return nil

	//连接服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial-错误，err= ", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// loginMes.UserName = userName
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes)-错误，err= ", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes)-错误，err= ", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn)-错误，err= ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
		//在这里启动一个协程，保证客户端和服务器的通讯
		//如有数据,接收并显示在客户端
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else { // if loginResMes.Code == 500
		fmt.Println(loginResMes.Error)
	}
	return

}

func (userPro *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial()-错误err=", err)
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal()-错误err=", err)
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal()-错误err=", err)
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg()-错误err=", err)
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("tf.ReadPkg()-错误err=", err)
	}
	//处理服务器发送过来的消息
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("恭喜你注册成功，现在赶紧登录吧！")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
