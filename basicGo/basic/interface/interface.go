package main

import (
	"fmt"
)

/*
	接口是一种类型，一种抽象的类型
*/

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct {

}


func (c cat) say() {
	fmt.Println("喵喵喵")
}

type person struct {
	name string
}

func(p person) say(){
	fmt.Printf("啊啊啊\n")
}

// 打的函数
func da(arg sayer) {
	arg.say()
}

/*
接口不管你是什么类型，它只管你要实现什么方法
定义一个类型，一个抽象的类型，只要实现了say()这个方法的类型都可以称为sayer的类型
 */
type sayer interface {
	say()
}


// 使用值接受者实现接口和使用指针接受者实现接口的区别
type mover interface {
	move()
}

// 接口的嵌套
type animal interface {
	sayer
	mover
}

type people struct {
	name string
	age int8
}

// 使用值接受者实现接口：类型的值和类型的指针都能保存到接口变量中
func (p people) move() {
	fmt.Printf("%s在跑...\n",p.name)
}

// 使用指针接受者实现接口：

// 任意类型都实现了空接口-- 空接口变量可以存储任意值
// 空接口一般不需要提前定义
type xx interface {

}

func main()  {
	c1 := cat{}
	da(c1)
	d1:= dog{}
	da(d1)
	p1:= person{
		name:"哪吒",
	}
	da(p1)

	var m mover
	pp := people{ // p1作为people 类型的值
		name: "小姚子",
		age:  18,
	}
	
	p10 := &people{ // p2是people类型的指针
		name: "哪吒",
		age:  20,
	}


	
	m = pp
	m =p10
	m.move()
	fmt.Println(m)

	var x interface{} // 定义一个空接口变量x
	x = "hello"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)

	var mm = make(map[string]interface{},16)
	mm["name"] = "娜娜"
	mm["sex"] = "girl"
	mm["age"] = 20
	mm["hobby"] = []string{"排球","羽毛球"}
	fmt.Println(mm)

	var xy interface{}
	 xy = "hello"
	 // xy = 125

	ret,ok:=xy.(string)	// 类型推断
	if !ok{
		fmt.Println("不是字符串")
	}else {
		fmt.Println("是字符串类型",ret)
	}
}

