package main

import (
	"fmt"
	"net"
)

/*
udp协议，通常做视频直播，会议等
和tcp不同，udp是无连接的传输层协议
udp server demo
 */
func main()  {
	listen,err:=net.ListenUDP("udp",&net.UDPAddr{
		IP:net.IPv4(127,0,0,1),
		Port:30000,
	})
	if err != nil{
		fmt.Printf("listen failed:%\n",err)
		return
	}
	defer listen.Close()
	fmt.Println("服务端已连接")

	for{
		var buf [1024]byte
		n,addr,err:= listen.ReadFromUDP(buf[:])
		if err != nil{
			fmt.Printf("read from udp failed,err:%v\n",err)
		}
		fmt.Println("接受到数据：",string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil{
			fmt.Printf("write to udp failed,err:%v\n",err,addr)
			return
		}

	}
}
