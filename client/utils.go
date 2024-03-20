package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/28267/chatroom/common/message"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Printf("\n服务器在等待接收来自客户端%v的消息...\n", conn.RemoteAddr().String())
	_, err = conn.Read(buf[:4])
	if err != nil {
		// fmt.Println("conn.Read-错误: ",err)
		return
	}
	// fmt.Println(buf[:n])
	// var pkgLen uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[:4])
	con, err := conn.Read(buf[:pkgLen])
	if con != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		// fmt.Println("json.Unmarshal(buf[:pkgLen],&mes)-错误: ",err)
		return
	}
	return

}
func writePkg(conn net.Conn, data []byte) (err error) {
	//将数据的长度转换成一个byte切片
	// var pkgLen uint32
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	con, err := conn.Write(buf[:4])
	if con != 4 || err != nil {
		fmt.Println("conn.Write-错误，err= ", err)
		return
	}
	// fmt.Printf("客户端发送了%d个字节的消息,内容是：%v",len(data),string(data))
	con, err = conn.Write(data)
	if con != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data)-错误，err= ", err)
		return
	}

	return
}
