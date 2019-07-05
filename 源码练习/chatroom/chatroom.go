package main

import (
	"fmt"
	"net"
)

//创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

//创建全局map,存储在线用户
var onlineMap map[string]Client
//创建全局channel,用于传递用户消息
var message = make(chan string)

func Manager() {
	//初始化 onlineMap
	onlineMap = make(map[string]Client)
	//监听全局channel中是否有数据，有数据存储至msg,无数据阻塞
	for {
		msg := <-message
		//循环发送消息给所有在线用户
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}

func handleConect(conn net.Conn) {
	defer conn.Close()
	//获取用户网络地址
	net_addr := conn.RemoteAddr().String()
	//创建新连接用户的结构体 默认用户名是Ip+Port
	clnt := Client{make(chan string), net_addr, net_addr}
	//将新连接的用户加入在线用户map，key:addr,value:用户结构体
	onlineMap[net_addr] = clnt

	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(clnt,conn)
	//发送用户上线消息到全局message通道中
	message <- clnt.Name + "login"
	for{
		;
	}
}

func WriteMsgToClient(clnt Client, conn net.Conn) {
	//监听用户自带channel是否有消息，有的话写给客户端
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listener.Close()

	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}
		//启动go程处理客户端数据请求
		go handleConect(conn)
	}
}
