package main

import (
	"fmt"
	"time"
)

func main(){
	// chan1()
	/*
	ch1:=make(chan int,100)
	ch2:=make(chan int,200)

	go f1(ch1)
	go f2(ch1,ch2)
	// 从通道中取值的方式
	for ret:=range ch2{
		fmt.Println(ret,"ret")
	}
	 */

	/*
	jobs:=make(chan int,100)
	results:=make(chan int,100)

	// 开启3个goroutine
	for j:=0;j<3;j++{
		go worker(j,jobs,results)
	}

	// 发送5个任务
	for i:=0;i<5;i++{
		jobs<-i
	}
	close(jobs)
	// 输出结果
	for i:=0;i<5;i++{
		ret:=<-results
		fmt.Println(ret)
	}
	*/

	ch:=make(chan int,2)
	for i:=0;i<10;i++{
		select {
			case x:= <-ch:
				fmt.Println(x)
			case ch<-i:
			default:
				fmt.Println("oop")
		}
	}
}

func worker(id int,jobs<-chan int,results chan<- int){
	for job:=range jobs{
		fmt.Println("--",id,job)
		results <- job*2
		time.Sleep(time.Second)
		fmt.Println("++",id,job)
	}
}

func chan1(){
	//	带缓冲区通道
	ch1 := make(chan int,3) // make()中的第二个参数是缓冲区大小
	ch1 <-10
	x:=<-ch1
	fmt.Println(x,"xx")
	close(ch1)
}

/*
生成0-100的数字发送到ch1
chan<- 表示只能发送值
*/
func f1(ch chan<- int)  {
	for i:=0;i<100 ;i++  {
		ch<-i
	}
	close(ch)
}

/*
从ch1中取出数据算它的平方，把结构发送到ch2中
<-chan 表示只能取值
*/
func f2(ch1 <-chan int,ch2 chan<- int){
	for{
		tmp,ok:=<-ch1
		if !ok{
			break
		}
		ch2<-tmp*tmp
	}
	close(ch2)
}