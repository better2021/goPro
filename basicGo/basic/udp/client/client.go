package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client demo
func main()  {
	c,err:=net.DialUDP("udp",nil, &net.UDPAddr{
		IP:  net.IPv4(127,0,0,1),
		Port: 30000,
		Zone: "",
	})

	if err !=nil{
		fmt.Printf("dial failed,err:%v\n",err)
		return
	}
	defer c.Close()
	fmt.Println("客户端已连接")

	input:= bufio.NewReader(os.Stdin)
	// for循环是为了不停地读写数据
	for{
		s,_:=input.ReadString('\n') // 通过换行符去读取数据
		_,err=c.Write([]byte(s)) // 网络发送的都是字节类型的数据 []byte(s)
		if err!=nil{
			fmt.Printf("send to server failed,err:%v\n",err)
			return
		}

		//  接受数据
		var buf [1024]byte
		n,addr,err:= c.ReadFromUDP(buf[:])
		if err!=nil{
			fmt.Printf("recv from udp failed,err:%v\n",err)
			return
		}
		fmt.Printf("read from %v,mag:%v\n",addr,string(buf[:n])) // string(buf[:n])把接受到的数据转换成字符串类型
	}
}