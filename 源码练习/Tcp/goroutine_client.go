package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dail error:", err)
		return
	}
	defer conn.Close()

	//获取用户键盘输入，将输入发送发送给服务器
	go getInput(conn)

	//回显服务器回发的大写数据
	buf:=make([]byte,4096)
	for{
		n, err := os.Stdin.Read(buf)
		if n==0 {
			fmt.Println("服务器检测到服务器关闭，关闭连接")
			return
		}
		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}
		fmt.Println("客户端读到服务器回显数据",string(buf[:n]))
	}
}
func getInput(conn net.Conn) {
	str := make([]byte, 4096)
	for{
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("os.stdin.Read error:", err)
			continue
		}
		conn.Write(str[:n])
	}
}
