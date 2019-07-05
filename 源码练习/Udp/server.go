package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//组织一个udp地址结构，指定服务器的IP和port
	serv_addr,err:=net.ResolveUDPAddr("udp","127.0.0.1:8001")
	if err !=nil{
		fmt.Println("net.ResolveUDPAddr error:",err)
		return
	}
	fmt.Println("UDP服务器地址结构创建完成！！！")
	//创建用户通信的socket
	udp_conn,err:=net.ListenUDP("udp",serv_addr)
	if err !=nil{
		fmt.Println("net.ListenUDP error:",err)
		return
	}
	defer udp_conn.Close()
	fmt.Println("UDP服务器通信socket创建完成！！！")
	buf :=make([]byte,4096)
	//返回三个值,分别是 读到的字节数，客户端的地址，error
	n,cli_addr,err:=udp_conn.ReadFromUDP(buf)
	if err !=nil{
		fmt.Println("ReadFromUDP error:",err)
		return
	}
	//模拟处理数据
	fmt.Printf("服务器读到：%v的数据:%s\n",cli_addr,string(buf[:n]))
	//回写数据给客户端
	daytime :=time.Now().String()
	_,err =udp_conn.WriteToUDP([]byte(daytime),cli_addr)
	if err !=nil{
		fmt.Println("WriteToUDP error:",err)
		return
	}
}
