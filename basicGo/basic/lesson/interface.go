package lesson

import "fmt"

/*
接口是一种类型，一种抽象的类型

*/

type dog struct {

}

type cat struct {

}

type people struct {
	name string
}

// 定义一个类型，抽象的类型，只要实现了say()这个方法的类型
type sayer interface {
	say()
}

type mover interface {
	move()
}

type animal interface {
	move()
	say()
}

func (d dog) say() {
	fmt.Println("汪汪")
}

func (c cat) say(){
	fmt.Println("喵喵")
}

func (p people) say() {
	fmt.Println("啊啊啊")
}

func (p people) move() {
	fmt.Println("人在跑")
}

func da(arg sayer)  {
	arg.say() // 调用方法
}

func Inter()  {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)

	p1:= people{name:"coco"}
	da(p1)

	p2 := &people{name:"qq"}

	var m mover
	m = p1
	m = p2
	m.move()

	var a animal
	a = p2
	fmt.Println("--")
	a.move()
	a.say()

	// 空接口
	var x interface{}
	x = "hello"
	x = 120
	x = false
	x = p2
	fmt.Println(x)
}