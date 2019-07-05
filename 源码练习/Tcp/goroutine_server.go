package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen error:", err)
		return
	}
	defer listener.Close()

	//监听客户端连接请求
	for {
		fmt.Println("服务器等待连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}
		defer conn.Close()

		//具体完成服务器和客户端的数据通信
		go connectHandler(conn)
	}
}
func connectHandler(conn net.Conn) {
	//获取连接的客户端的Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, "成功连接")

	//循环读取客户端发送的数据
	buf :=make([]byte,4096)//创建缓冲区
	for{
		n,err:=conn.Read(buf)
		if n==0 {
			fmt.Println("服务器检测到客户端关闭，关闭连接")
			return
		}
		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务器检测到客户端退出请求，关闭连接")
			return
		}
		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}
		fmt.Println("服务器读到数据：",string(buf[:n]))

		//完成小写转大写，回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
