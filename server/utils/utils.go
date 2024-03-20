package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/28267/chatroom/common/message"

	// "io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {

	// buf := make([]byte, 8096)
	fmt.Printf("\n服务器在等待接收来自客户端%v的消息...\n", transfer.Conn.RemoteAddr().String())
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read-错误: ",err)
		return
	}
	// fmt.Println(buf[:n])
	// var pkgLen uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(transfer.Buf[:4])
	con, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if con != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		// fmt.Println("json.Unmarshal(buf[:pkgLen],&mes)-错误: ",err)
		return
	}
	return

}

func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//将数据的长度转换成一个byte切片
	// var pkgLen uint32
	var pkgLen uint32 = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(transfer.Buf[:4], pkgLen)

	con, err := transfer.Conn.Write(transfer.Buf[:4])
	if con != 4 || err != nil {
		fmt.Println("transfer.conn.Write-错误，err= ", err)
		return
	}
	fmt.Printf("\n客户端发送了%d个字节的消息,内容是：%v\n", len(data), string(data))
	con, err = transfer.Conn.Write(data)
	if con != int(pkgLen) || err != nil {
		fmt.Println("transfer.conn.Write(data)-错误，err= ", err)
		return
	}

	return
}
