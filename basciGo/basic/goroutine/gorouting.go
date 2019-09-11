package main

import (
	"fmt"
	"runtime"
	"sync"
)

// sync.WaitGroup异步等待
var wg sync.WaitGroup

func main(){
	/*
	wg.Add(1000) // 计数+1
	for i:=0;i<1000;i++{
		// go hello(i) // 开启一个goroutine去执行hello函数
		go func(i int) {
			fmt.Println("go",i) // 利用闭包特性把i传进来可以打印出1-1000
		}(i)
	}

	fmt.Println("func main")
	wg.Wait() // 等所以小弟都干完活之后收兵
	 */

	runtime.GOMAXPROCS(1) // 只利用CPU的单核
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	// time.Sleep(time.Second)
}

func hello(i int)  {
	fmt.Println("hello go",i)
	wg.Done() // 通知wg把计数器-1
}

func a()  {
	for i:=0;i<10;i++  {
		fmt.Println(i,"A")
	}
	wg.Done()
}

func b()  {
	for i:=0;i<10;i++  {
		fmt.Println(i,"B")
	}
	wg.Done()
}