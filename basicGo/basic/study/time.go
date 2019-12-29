package study

import (
	"fmt"
	"reflect"
	"time"
)

func Time(){
	//createTime()
	//parseTime()
	//getTime()
	stamTime()
}


// 创建时间
func createTime(){
	// 当前时间
	nowtime := time.Now()
	fmt.Printf("当前时间为：%v\n",nowtime)

	// 自定义时间
	custime := time.Date(2008,7,15,16,30,20,0,time.Local)
	fmt.Println(custime)
}

// 时间格式化与解析
func parseTime(){
	nowTime := time.Now()
	stringtime := nowTime.Format("2006年01月02日 15:04")
	fmt.Println(stringtime)

	// 时间解析
	strTime := "2019-01-10 15:20:10"
	objTime,_ := time.Parse("2006-01-02 15:04:05",strTime)
	fmt.Println(objTime)
}

// 获取年月日
func getTime(){
	nowTime := time.Now()
	year,month,day := nowTime.Date()
	fmt.Println(year,month,day)

	hour,min,sec := nowTime.Clock()
	fmt.Println(hour,min,sec)

	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Day())
	fmt.Println(nowTime.Hour())

	fmt.Println(nowTime.Weekday())

	time :=  nowTime.YearDay()

	fmt.Println("今年已经过了",time,"天") // 指今年一共过了多少天
}

// 时间戳
func stamTime(){
	nowTime := time.Now()
	fmt.Println(nowTime.Unix()) // 时间戳(单位秒)

	// 时间间隔
	fmt.Println(nowTime.Add(time.Minute*10)) // 10分钟后
	fmt.Println(nowTime.AddDate(1, 0, 0))	// 1年后

	time.Sleep(time.Second*3) //  程序睡眠3秒钟

	/*----------*/
	fmt.Println(reflect.TypeOf(nowTime)) // 通过反射获取变量的类型
	fmt.Println(reflect.ValueOf(nowTime)) // 通过反射获取变量的值

	type Person struct {
		Name string
		Age int
	}
	p := Person{ "lisi", 13}
	fmt.Println(reflect.TypeOf(p).Name())
	fmt.Println(reflect.TypeOf(p).Kind())

	str := p
	fmt.Println(reflect.TypeOf(str).Name())

	switch reflect.TypeOf(str).Name() {
		case "string":
			fmt.Println("字符串")
		case "int":
			fmt.Println("数字")
		case "bool":
			fmt.Println("布尔值")
		default:
			fmt.Println("其他类型")
		}
}