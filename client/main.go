package main

import (
	"fmt"
	"os"
)

func main() {
	var userId int
	var userPwd string
	var key int
	var loop bool = true

	for loop {
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
			// loop = false
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出聊天系统")
			// loop = false
			os.Exit(0)
		}
		if key == 1 {
			fmt.Println("请输入用户id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			Login(userId, userPwd)
		} else if key == 2 {
			fmt.Println("进行用户注册")
		}
	}

}
