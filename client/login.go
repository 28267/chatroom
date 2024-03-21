package main

import (
	"encoding/json"
	"fmt"
	"net"
	_ "time"

	"github.com/28267/chatroom/common/message"
)

func Login(userId int, userPwd string) (err error) {

	// fmt.Printf("userId=%v,userPwd=%v",userId,userPwd)
	// return nil

	//连接服务器
	conn, err := net.Dial("tcp", "192.168.1.5:8888")
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
	err = writePkg(conn, data)

	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn)-错误，err= ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return

}
