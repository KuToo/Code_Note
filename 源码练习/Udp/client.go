package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器的协议，Ip以及端口号 通信套接字
	conn ,err := net.Dial("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.dail error:", err)
		return
	}
	defer conn.Close()

	//主动截数据给服务器
	conn.Write([]byte("Are You Ready ?"))
	buf :=make([]byte,4096)
	n,err :=conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	fmt.Println("服务器回发：",string(buf[:n]))
}
