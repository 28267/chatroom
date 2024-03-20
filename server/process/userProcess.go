package process2

import (
	// "encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/28267/chatroom/common/message"
	"github.com/28267/chatroom/server/model"
	"github.com/28267/chatroom/server/utils"

	// "io"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (userProcess *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType //消息类型
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),loginMes)-错误: ", err)
		return
	}
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
	}

	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 200 //合法登录
	// } else {
	// 	loginResMes.Code = 500 //不合法登录
	// 	loginResMes.Error = "该用户不存在，请注册再使用"
	// }

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes)-错误: ", err)
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes)-错误: ", err)
	}
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	// fmt.Printf("客户端发送了 %v 字节的数据\n", con)
	return
}

func (userProcess *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal()-错误: ", err)
	}
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.marshal()-错误: ", err)
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal()-错误: ", err)
	}
	tf := utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}
