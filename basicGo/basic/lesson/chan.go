package lesson

import (
	"fmt"
	"time"
)

func Ch()  {
	/*
	非缓冲通道：make(chan T)
		 一次发送，一次接受，都是阻塞的

	缓冲通道：make(chan T,capacity)
		发送：缓冲区的数据满了，才会阻塞
		接受：缓冲区的数空了，才会阻塞
	*/

	ch1 := make(chan int)	// 非缓冲通道
	fmt.Println(len(ch1),cap(ch1))
	time.Sleep(time.Second)
	//ch1 <- 100 // 阻塞式，需要有其他的goroutine解除阻
	//data := <- ch1
	//fmt.Println(data)

	ch2 := make(chan int,5) // 缓冲通道，缓冲区大小为5
	fmt.Println(len(ch2),cap(ch2))

	ch2 <- 150
	ch2 <- 200
	ch2 <- 500
	fmt.Println(len(ch2),cap(ch2))
}