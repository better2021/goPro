package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client demo
func main()  {
	// 1.与服务端建立连接
	conn,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("dial failed",err)
		return
	}
	fmt.Println("客户端已连接")
	// 2.利用该连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for{
		s,_:=input.ReadString('\n')
		str := strings.TrimSpace(s) // 去除字符串空格
		if strings.ToUpper(str)=="Q"{ // 输入大写Q退出消息发送
			return
		}
		// 给服务器发消息
		_,err:=conn.Write([]byte(str))
		if err !=nil{
			fmt.Println("send failed" ,err)
			return
		}
		// 从服务端接受恢复的消息
		var buf [1024]byte
		n,err:=conn.Read(buf[:])
		if err !=nil{
			fmt.Println("read failed",err)
			return
		}
		fmt.Println("收到服务端回复：",string(buf[:n]))
	}
}