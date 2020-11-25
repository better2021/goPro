package main

import (
	"fmt"
	"time"
)

/*
标准库time
*/
func main(){
	now:= time.Now()
	fmt.Println(now,"now")
	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	hour:= now.Hour()
	min:= now.Minute()
	second:= now.Second()
	fmt.Println(year,month,day,hour,min,second)
	timeStamp1:= now.Unix() // 单位：秒
	timeStamp2:=now.UnixNano()	// 单位：纳秒
	fmt.Println(timeStamp1,timeStamp2)

	// 将时间戳转换为具体的时间格式
	// 1568084513 + 3600
	t:=time.Unix(1568084513 + 3600,0)
	fmt.Println(t)
	// 时间间隔
	n:=3
	time.Sleep(time.Duration(n)*time.Second) // Duration类型
	fmt.Println("over")

	t2 := now.Add(2*time.Hour) // 时间相加，加2个小时
	fmt.Println(t2)
	// sub
	fmt.Println(t2.Sub(now)) //2g0m0s

	//  定时器
	//for tmp:=range time.Tick(2*time.Second){ // 2秒钟执行一次
	//	fmt.Println(tmp)
	//}

	// 时间格式化
	ret := now.Format("2006-01-02")
	reg := now.Format("2006.01.02 15:04:05")
	fmt.Println(ret,"ret")
	fmt.Println(reg,"reg")

	// 解析字符串类型的时间
	loc,err:=time.LoadLocation("Asia/Shanghai")
	if err!=nil{
		fmt.Println(err)
		return
	}
	// 根据时区去解析一个字符串格式的时间
	timeStr:="2018/08/07 15:00:00"
	timeObj,err:=time.ParseInLocation("2006/01/02 15:04:05",timeStr,loc)
	if err !=nil{
		fmt.Println(err,"err")
		return
	}
	fmt.Println(timeObj)
}
