package main

import (
	// "encoding/binary"
	// "encoding/json"
	"fmt"
	"time"

	"github.com/28267/chatroom/server/model"

	// "go_code/chatroom/common/message"
	// "io"
	"net"
)

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//将数据的长度转换成一个byte切片
// 	// var pkgLen uint32
// 	var pkgLen uint32 = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[:4], pkgLen)

// 	con, err := conn.Write(buf[:4])
// 	if con != 4 || err != nil {
// 		fmt.Println("conn.Write-错误，err= ", err)
// 		return
// 	}
// 	// fmt.Printf("客户端发送了%d个字节的消息,内容是：%v",len(data),string(data))
// 	con, err = conn.Write(data)
// 	if con != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(data)-错误，err= ", err)
// 		return
// 	}

// 	return
// }
// func readPkg(conn net.Conn) (mes *message.Message, err error) {

// 	buf := make([]byte, 8096)
// 	fmt.Printf("\n服务器在等待接收来自客户端%v的消息...\n", conn.RemoteAddr().String())
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		// fmt.Println("conn.Read-错误: ",err)
// 		return
// 	}
// 	// fmt.Println(buf[:n])
// 	// var pkgLen uint32
// 	var pkgLen uint32 = binary.BigEndian.Uint32(buf[:4])
// 	con, err := conn.Read(buf[:pkgLen])
// 	if con != int(pkgLen) || err != nil {
// 		return
// 	}
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		// fmt.Println("json.Unmarshal(buf[:pkgLen],&mes)-错误: ",err)
// 		return
// 	}
// 	return

// }
//
//	func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//		var loginMes message.LoginMes
//		err = json.Unmarshal([]byte(mes.Data), &loginMes)
//		if err != nil {
//			fmt.Println("json.Unmarshal([]byte(mes.Data),loginMes)-错误: ", err)
//			return
//		}
//		var resMes message.Message
//		resMes.Type = message.LoginResMesType //消息类型
//		var loginResMes message.LoginResMes
//		if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//			loginResMes.Code = 200 //合法登录
//		} else {
//			loginResMes.Code = 500 //不合法登录
//			loginResMes.Error = "该用户不存在，请注册再使用"
//		}
//		data, err := json.Marshal(loginResMes)
//		if err != nil {
//			fmt.Println("json.Marshal(loginResMes)-错误: ", err)
//		}
//		resMes.Data = string(data)
//		data, err = json.Marshal(resMes)
//		if err != nil {
//			fmt.Println("json.Marshal(resMes)-错误: ", err)
//		}
//		err = writePkg(conn, data)
//		// fmt.Printf("客户端发送了 %v 字节的数据\n", con)
//		return
//	}
//
//	func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//		switch mes.Type {
//		case message.LoginMesType:
//			err = serverProcessLogin(conn, mes)
//			//处理登录
//		// case message.RegisterMesType:
//		//处理注册
//		default:
//			fmt.Println("该消息类型不存在，无法处理")
//		}
//		return
//	}
func process(conn net.Conn) {
	defer conn.Close()
	// exitChan <- true
	// defer close(exitChan)
	for {
		processor := &Processor{
			Conn: conn,
		}
		err := processor.Process2()
		if err != nil {
			fmt.Println("客户端和服务器通讯协程错误=err", err)
			return
		}
	}

}
func InitUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	InitPool("localhost:6379", 16, 0, 300*time.Second)
	InitUserDao()
	fmt.Println("服务器（新的结构）在8888端口监听...")

	listen, err := net.Listen("tcp", "192.168.1.7:8888")
	if err != nil {
		fmt.Println("错误，err= ", err)
	}
	defer listen.Close()

	for {
		fmt.Println("服务器在等待客户端的连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("错误，err= ", err)
			return
		} else {
			fmt.Printf("客户端%v连接到服务器\n", conn.RemoteAddr().String())
		}

		// exitChan := make(chan bool, 1)
		go process(conn)
		// for {
		// 	_, ok := <-exitChan
		// 	if !ok {
		// 		break
		// 	}
		// }
	}

}
