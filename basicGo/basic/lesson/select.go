package lesson

import (
	"fmt"
	"time"
)

func Sele(){
	/*
	分支语句，if,switch,select
	select 语句类似于switch语句
	*/

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <-100
	}()

	go func() {
		time.Sleep(2*time.Second)
		ch2 <-200
	}()

	select {
		case num1 := <-ch1:
			fmt.Println("data from ch1",num1)
		case num2,ok := <-ch2:
			if ok{
				fmt.Println("data from ch2",num2)
			}else {
				fmt.Println("ch2 is close")
			}
		//default:
		//	fmt.Println("default...")
	}

	fmt.Println("over")
}