package study

import (
	"fmt"
	"runtime"
	"time"
)

/*
Go语言从语言层面原生提供了协程支持，即goroutine，Go语言内部实现了多个goroutine之间的内存共享，让并发编程变得极为简单。
执行goroutine只需极少的栈内存(大概是4~5KB)，可同时运行成千上万个并发任务。
*/
func Seven(){
	//go say("hello") // 以协程方式执行say函数
	//say("golang") // 以普通方式执行say函数
	//time.Sleep(3*time.Second) // 睡眠3秒：防止协程未执行完毕，主程序退出

	//go running()
	//var input string
	//fmt.Scanln(&input)

	//cpuNum := runtime.NumCPU()  //获取当前系统的CPU核心数
	//fmt.Println(cpuNum)
	//runtime.GOMAXPROCS(6) //Go中可以轻松控制使用核心数

	//sched()
	exit()
}

// goroutine由Go的runtime管理，通过关键字go实现，示例：
func say(s string){
	for i:=0;i<3;i++{
		fmt.Println(s)
	}
}


// 同时执行两件事
func running(){
	var times int
	for{
		times++
		fmt.Println("tick",times)
		time.Sleep(time.Second)
	}
}

// runtime.Gosched():用于让出CPU时间片，即让出当前协程的执行权限，调度器安排其他等待的任务运行。
func sched(){
	for i:=1;i<=10;i++{
		go func(i int) {
			if i== 5{
				runtime.Gosched() // 协程让出，5永远不会第一输出
			}
			fmt.Println(i) // 打印一组无规律数字
		}(i)
	}
	time.Sleep(time.Second)
}

// runtime.Goexit():用于立即终止当前协程运行，调度器会确保所有已注册defer延迟调用被执行。
func exit(){
	for i:=1;i<=5;i++{
		defer fmt.Println("defer",i)
		go func(i int) {
			if i==3{
				runtime.Goexit()
			}
			fmt.Println(i) // i==3时runtime.Goexit() 终止当前协程，defer延迟调用被执行
		}(i)
	}
	time.Sleep(time.Second)
}