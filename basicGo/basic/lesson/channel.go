package lesson

import (
	"fmt"
	"time"
)

func Chan()  {
	/*
	channel 通道
	data := <- a // 读取数据data从通道a中
	a <- data // 把data数据写入通道a
	*/
	//var a chan int
	//fmt.Printf("%T,%v\n",a,a)
	//
	//if a==nil{
	//	fmt.Println("先要创建通道...")
	//	a = make(chan int)
	//	fmt.Println(a)
	//}

	//test01()
	//test02()

	ch1 := make(chan int)
	go sendData(ch1)

	// 读取通道的数据
	for {
		time.Sleep(time.Second)

		v,ok := <-ch1
		if !ok{
			fmt.Println("已读取数据",ok)
			break
		}
		fmt.Println(v,ok)
	}

}

func test01()  {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println(" 子goroutine",i)
		}
		// 循环结束后向通道中写数据
		ch1 <- true
		fmt.Println("over")
	}()

	data := <-ch1	// 阻塞
	fmt.Println("main",data)
}

func test02()  {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine")
		data := <-ch1 // 从ch1中读取数据
		fmt.Println(data,"data")
	}()

	ch1 <- 10 // 写入数据10到通道ch1
	fmt.Println(ch1,"ch1")
}

func sendData(ch1 chan int){
	// 发送10条数据
	for i:= 0;i<10; i++{
		ch1 <- i // 将i写入通道ch1
	}

	close(ch1) // 关闭通道
}