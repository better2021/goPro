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

	chanDemo()
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
