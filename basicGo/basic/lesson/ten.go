package lesson

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex

func Ten()  {
	//go printNum()
	//time.Sleep(1*time.Second)
	//go printStr()
	//time.Sleep(time.Second)

	/*临界资源*/
	//a := 1
	//go func() {
	//	a = 2
	//	fmt.Println("goroutine",a)
	//}()
	//
	//a = 3
	//time.Sleep(time.Second) // sleep 后就会去执行goroutine （go func函数）
	//fmt.Println("main goroutine",a)

	/*
	4个goroutine，模拟4个售票口
	在使用互斥锁的时候，对资源操作完成，一定要解锁，不然会出现程序异常，死锁等问题
	*/

	//wg.Add(4)
	//go saleTicket("售票口1")
	//go saleTicket("售票口2")
	//go saleTicket("售票口3")
	//go saleTicket("售票口4")
	//
	//wg.Wait()

	//time.Sleep(5*time.Second)

	/*
	waitGroup 同步等待组
	Add() 设置等待组中要执行的子 goroutine的数量
	Wait() 让主goroutine处于等待
	*/
	//wg.Add(2)
	//go printNum()
	//go printStr()

	//fmt.Println("进入阻塞状态")
	//wg.Wait() // 表示goroutine进入等待，意味着阻塞
	//fmt.Println("解除阻塞")

	/*
	读写锁
	*/
	rwMutex = new(sync.RWMutex)

	wg.Add(3)
	go readData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("over")

}



var wg sync.WaitGroup // 创建一个同步等待组的对象

// 全局变量，表示票
var ticket = 10 // 10张票
var mutex sync.Mutex // 创建锁头
func saleTicket(name string){
	rand.Seed(time.Now().UnixNano())
	for{
		// 上锁
		mutex.Lock() // 上锁后就只能有一个goroutine来执行，防止多个goroutine之前抢占资源
		if ticket > 0{
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出",ticket)
			ticket--
		}else {
			mutex.Unlock() // 解锁
			fmt.Println(name,"售罄，没有票了...")
			break
		}
		mutex.Unlock() // 解锁
	}
	defer wg.Done()
}

func printNum()  {
	for i:=1;i<=100;i++{
		fmt.Println("打印数字",i)
	}

	 wg.Done() // 给wg等待组中的counter数值减1，同 wg.Add(-1)
}

func printStr(){
	defer wg.Done()
	for i:=1;i<=100;i++{
		fmt.Println("字母",i)
	}
}

// 读
func readData(i int) {
	fmt.Println(i,"开始读：read start...")
	rwMutex.RLock() // 读操作上锁
	fmt.Println(i,"正在读取数据：reading...")
	time.Sleep(time.Second)
	rwMutex.RUnlock() // 读操作解锁
	fmt.Println(i,"读结束...")
	defer wg.Done()
}

// 写
func writeData(i int){
	defer wg.Done()
	fmt.Println(i,"开始写...")
	rwMutex.Lock()
	fmt.Println(i,"正在写...")
	time.Sleep(time.Second)
	rwMutex.Unlock()
	fmt.Println(i,"写结束")
}