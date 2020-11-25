package lesson

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func Time()  {
	//mapFun()
	inter()
}

func timeStr(){
	d := time.Now()
	fmt.Println(d)

	local,err := time.LoadLocation("Asia/Shanghai")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(local)

	year, month, day := d.Date()
	fmt.Println(year,month,day)

	hour, minute, second := d.Clock()
	fmt.Println(hour, minute, second)

	fmt.Println(d.Year(),d.Month())
	fmt.Println(d.Weekday())
	fmt.Println(d.Unix())
	fmt.Println(d.Local())
	fmt.Println(d.Location())

	fmt.Println(strconv.Atoi("123")) // 字符串转整形 string => int
	fmt.Println(strconv.Itoa(123)) // 整形转字符串


	// 以xx结尾的字节,返回true或false
	fmt.Println(bytes.HasPrefix([]byte("Gopher"),[]byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"),[]byte("er")))

	// 以xx结尾的字节,返回true或false
	fmt.Println(bytes.HasSuffix([]byte("Gopher"),[]byte("Go")))
	fmt.Println(bytes.HasSuffix([]byte("Gopher"),[]byte("er")))

	fmt.Println(len("12as"))
}

func mapFun(){
	/*
	存储的是无序的键值对
	剑不能重复，并且和value值意义对应的
	*/
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"GO":98,"Python":84,"java":65,"js":99}
	map3["Python"] = 90
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	v1,ok := map3["java"]
	if ok{
		fmt.Println(v1)
	}else {
		fmt.Println("不存在java这个key")
	}

	// 删除数据
	delete(map3,"java")
	fmt.Println(map3)

}


// 定义空接口
//type A interface {}
//
//type Cat struct {
//	Color string
//}
//
//type Person struct {
//	Name string
//	Age int
//}
//
//func test1(a A){
//	fmt.Println(a)
//}

func inter()  {
	//var a1 A = Cat{Color:"花猫"}
	//var a2 A = person{Name:"yaozi",Age:18}
	//
	//fmt.Println(a1)
	//fmt.Println(a2)
	//test1(a1)

	//var cat Cat = Cat{}
	//cat.test1()
	//cat.test2()
	//cat.test3()
}

/**
	接口的前台
*/
// type A interface {
// 	test1()
// }
//
// type B interface {
// 	test2()
// }
//
// type C interface {
// 	A
// 	B
// 	test3()
// }
//
// type Cat struct {}
//
// func (c Cat) test1(){
// 	fmt.Println("test1")
//}
//
//func (c Cat) test2() {
//	fmt.Println("test2")
//}
//
//func (c Cat) test3() {
//	fmt.Println("test3")
//}


/*
 接口断言
方式一：
1.instance := 接口对象.(实际类型) // 不安全，会panic()
2.instance,ok := 接口对象.(实际类型) // 安全
方式二：switch
switch instance := 接口对象.(type){
	case 实际类型1：
		...
	case 实际类型2：
		...
}
*/

// 1.定义一个接口
type Shape interface {
	peri() float64  // 周长
	area() float64	// 面基
}

// 2.定义一个实现类：三角形
type Teiangle struct {
	a,b,c float64
}

