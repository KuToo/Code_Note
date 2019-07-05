package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//获取命令行参数
	list:=os.Args
	if len(list)!=2{
		fmt.Println("格式为：go run xxx.go 文件绝对路径")
		return
	}
	//提取文件绝对路径
	file_path :=list[1]
	file_info,err:=os.Stat(file_path)
	//偷取文件名
	if err!=nil {
		fmt.Println("os.Stat error",err)
		return
	}
	file_name :=file_info.Name()

	//主动发起连接请求
	conn,err:=net.Dial("tcp","127.0.0.1:8008")
	if err!=nil {
		fmt.Println("net.Dial error",err)
		return
	}
	defer conn.Close()

	//发送文件名给接收端
	conn.Write([]byte(file_name))

	//读取服务器回执 ok
	buf :=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println("conn.Read error",err)
		return
	}

	if "ok" == string(buf[:n]) {
		//写文件内容给服务器 ===借助网络
		sendFile(conn,file_path)
	}
}
func sendFile(conn net.Conn,file_path string){
	f,err :=os.Open(file_path)
	if err!=nil {
		fmt.Println("os.Open error",err)
		return
	}
	defer f.Close()
	//从文件中读数据，写给网络接收端，读多少，写多少
	buf:=make([]byte,4096)
	for{
		n,err:=f.Read(buf)
		if err!=nil {
			if err==io.EOF {
				fmt.Println("发送文件完成")
			}else{
				fmt.Println("f.Read error",err)
			}
			return
		}

		//写到网络socket中
		_,err=conn.Write(buf[:n])
		if err!=nil {
			fmt.Println("conn.Write error",err)
			return
		}



	}
}