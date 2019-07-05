package main

import (
	"fmt"
	"net"
	"strings"
	"time"
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
	//创建 判断用户是否活跃 的通道
	is_active :=make(chan bool)
	//获取用户网络地址
	net_addr := conn.RemoteAddr().String()
	//创建新连接用户的结构体 默认用户名是Ip+Port
	clnt := Client{make(chan string), net_addr, net_addr}
	//将新连接的用户加入在线用户map，key:addr,value:用户结构体
	onlineMap[net_addr] = clnt

	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(clnt,conn)
	//发送用户上线消息到全局message通道中
	message <- makeMsg(clnt,"login")

	//创建一个channel,用来判断用户退出状态
	is_quit:=make(chan bool)

	//创建一个匿名go程，专门处理用户发送的消息
	go func(){
		buf:=make([]byte,4096)
		for{
			n,err:=conn.Read([]byte(buf))
			//客户端退出
			if n==0 {
				is_quit<-true
				fmt.Printf("客户端%s退出\n",clnt.Name)
				return
			}
			if err!=nil {
				fmt.Println("conn.Read error",err)
				return
			}
			//将读到的用户消息报讯到msg中，string类型
			msg:=string(buf[:n-1])
			//提取用户在线列表 who命令
			if msg=="who" && len(msg)==3 {
				conn.Write([]byte("online user list:\n"))
				//遍历当前map，获取在线用户
				for _, user := range onlineMap {
					userinfo := user.Name + "\n"
					conn.Write([]byte(userinfo))
				}
			//修改自己的用户名
			}else if msg[:7]=="rename|" && len(msg)>=8{
				new_name:=strings.Split(msg,"|")[1]
				clnt.Name=new_name//修改结构体Name
				onlineMap[net_addr]=clnt
				conn.Write([]byte("rename successful\n"))
			}else{
				message<-makeMsg(clnt,msg)
			}
			is_active<-true
		}
	}()
	//保证不退出
	for{
		select {
			case <-is_quit://客户端退出
				//将用户从在线用户列表中移除
				delete(onlineMap,clnt.Addr)
				//向其他用户广播退出消息
				message<-makeMsg(clnt,"logout")
				return
			case <-is_active://客户端是否活跃
				//什么都不做，目的是重置下面的case倒计时
			case <-time.After(time.Second*10)://客户端退出
				//将用户从在线用户列表中移除
				delete(onlineMap,clnt.Addr)
				//向其他用户广播退出消息
				message<-makeMsg(clnt,"timeout logout")
				return
		}
	}
}

func WriteMsgToClient(clnt Client, conn net.Conn) {
	//监听用户自带channel是否有消息，有的话写给客户端
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func makeMsg(clnt Client,msg string) (buf string){
	buf ="["+clnt.Addr+"]"+clnt.Name+":"+msg
	return
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
