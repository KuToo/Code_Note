package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//创建用于监听的socket
	listener,err:=net.Listen("tcp","127.0.0.1:8008")
	if err!=nil {
		fmt.Println("net.Listen error",err)
		return
	}
	defer listener.Close()

	//阻塞监听
	conn,err:=listener.Accept()
	if err!=nil {
		fmt.Println("listener.Accept error",err)
		return
	}
	defer conn.Close()

	//获取文件名，保存
	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println("conn.Read error",err)
		return
	}
	//回写ok给发送端
	conn.Write([]byte("ok"))

	//获取文件内容
	file_name :=string(buf[:n])
	recvFile(conn,file_name)
}

func recvFile(conn net.Conn,file_name string){
	f,err:=os.Create(file_name)
	if err!=nil {
		fmt.Println("os.Create error",err)
		return
	}
	defer f.Close()

	//从网络中读取数据，写入本地文件
	buf:=make([]byte,4096)
	for{
		n,_:=conn.Read(buf)
		if n==0 {
			fmt.Println("接收文件完成")
			return
		}
		f.Write(buf[:n])
	}
}