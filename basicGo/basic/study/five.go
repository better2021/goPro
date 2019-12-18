package study

import (
	"fmt"
	"time"
)

func Five(){
	//for i:=0;i<100;i++{
	//	go func(i int) {
	//		for {
	//			fmt.Printf("Hello from" + "goroutine %d",i)
	//		}
	//	}(i)
	//}
	// time.Sleep(time.Millisecond)

	//chanDemo()

	/*
	iota 是go的关键字出现时将被重置为0，const中每新增一行常量声明将使iota计数一次
	iota 实行枚举类型
	*/

	const (
		n1 = iota // 0
		n2 = 10
		n3 = iota
		n4
		n5
	)

	fmt.Println(n2,n4,n5)
	fmt.Println(n1+n3)
}

func chanDemo(){
	c := make(chan int)
	go func() {
		n := <-c
		fmt.Println(n)
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
