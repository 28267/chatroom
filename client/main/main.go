package main

import (
	"fmt"
	"os"

	"github.com/28267/chatroom/client/process"
)

func main() {
	var userId int
	var userPwd string
	var userName string
	var key int
	// var loop bool = true

	for {
		fmt.Println("--------------------欢迎来到多人聊天室--------------------")
		fmt.Println("                    1.登录聊天室")
		fmt.Println("                    2.注册用户")
		fmt.Println("                    3.退出系统")
		fmt.Println()
		fmt.Println("请选择1~3")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			up := process.UserProcess{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户名称")
			fmt.Scanln(&userName)
			up := process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出聊天系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

}
