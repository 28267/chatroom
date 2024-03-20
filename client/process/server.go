package process

import (
	"fmt"
	"net"
	"os"

	"github.com/28267/chatroom/client/utils"
)

func ShowMenu() {
	fmt.Println("--------恭喜XXX登陆成功---------")
	fmt.Println("--------1.显示在线用户列表---------")
	fmt.Println("--------2.发送消息---------")
	fmt.Println("--------3.信息列表---------")
	fmt.Println("--------4.退出系统---------")
	fmt.Println()
	fmt.Println("请选择1~4")
	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("--------显示在线用户列表---------")
	case 2:
		fmt.Println("--------发送消息---------")
	case 3:
		fmt.Println("--------信息列表---------")
	case 4:
		fmt.Println("--------你选择了退出系统---------")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确..")
	}
}
func ServerProcessMes(conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息..")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg(conn)-错误，err= ", err)
			return
		}
		//客户端读到消息，处理下一步
		fmt.Println(mes)
	}

}
