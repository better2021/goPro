package lesson

import (
	"fmt"
)

func Ch()  {
	/*
	非缓冲通道：make(chan T)
		 一次发送，一次接受，都是阻塞的

	缓冲通道：make(chan T,capacity)
		发送：缓冲区的数据满了，才会阻塞
		接受：缓冲区的数空了，才会阻塞
	*/

	//ch1 := make(chan int)	// 非缓冲通道
	//fmt.Println(len(ch1),cap(ch1))
	//time.Sleep(time.Second)
	//ch1 <- 100 // 阻塞式，需要有其他的goroutine解除阻
	//data := <- ch1
	//fmt.Println(data)

	//ch2 := make(chan int,5) // 缓冲通道，缓冲区大小为5
	//fmt.Println(len(ch2),cap(ch2))
	//
	//ch2 <- 150
	//ch2 <- 200
	//ch2 <- 500
	//fmt.Println(len(ch2),cap(ch2))

	/*
	单向(定向)
	chan <- T,只支持写
	<- chan T,只读
	*/

	ch1 := make(chan int) // 双向，读，写
	//ch2 := make(chan <- int) // 单向，只能写，不能读
	//ch3 := make(<- chan int) // 单向，只能读，不能写

	go fun1(ch1) // 可读，可写

	data := <- ch1
	fmt.Println(data)


	go fun2(ch1)
	ch1 <- 200

}


// 该函数，只能操作只写的通道
func fun1(ch chan <- int)  {
	ch <- 100
	fmt.Println("fun1")
}

// 只读通道
 func fun2(ch <- chan int){
 	data :=<- ch
 	fmt.Println("fun2",data)
 }