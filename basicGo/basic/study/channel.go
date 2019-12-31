package study

import (
	"fmt"
	"time"
)

func Chan(){
	// test()
	//ran()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go fn1(ch1)
	go fn2(ch2)

	// select用于监听channel上的数据流动，在有多个channel时，不会让其串行执行，在上一章中，可以使用select让获取ch1直接3秒钟内完成，获取ch2在6秒内完成
	select {
		case r1 := <- ch1:
			fmt.Println("r1",r1)
		case r2 := <- ch2:
			fmt.Println("r2",r2)
		}
}

func test(){
	var ch chan int // 声明一个channel，但未初始化，值为nil
	ch = make(chan int) // 初始化，此时ch有了地址

	// 协程中向通道填充数据
	go func(){
		ch <- 100
		fmt.Println(ch)
		ch <- 200
		fmt.Println(ch)
		ch <- 300
		fmt.Println(ch)
	}()
	time.Sleep(time.Second)

	// 主协程中取出数据
	data1 := <-ch
	data2 := <-ch
	fmt.Println(data1,data2)
}

func ran(){
	//ch := make(chan int)
	//go func() {
	//	fmt.Println("start")
	//	ch <- 0
	//	fmt.Println("end")
	//}()
	//
	//<- ch
	//fmt.Println("all done")

	ch := make(chan int ,3)
	fmt.Println("缓冲长度",len(ch))
	ch <- 1
	ch <- 2
	fmt.Println("缓冲长度",len(ch))
}

func fn1(ch chan string){
	time.Sleep(time.Second*3)
	ch <- "fn11"
}

func fn2(ch chan string)  {
	time.Sleep(time.Second*3)
	ch <- "fn222"
}