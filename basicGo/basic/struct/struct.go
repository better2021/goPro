package main

import (
	"fmt"
)

// 定义结构体

type person struct {
	name,city string // 类型相同可以简写
	age int8
	info // 嵌套结构体
}

type  info struct {
	sex string
	email string
}

// 结构体指针
type animal struct {
	name string
	age int8
}

func main(){
	var p person
	p.name ="哪吒"
	p.age = 18
	p.city = "深圳"

	fmt.Println(p)
	fmt.Println(p.name)

	// 匿名结构体
	var user struct {
		name string
		married bool
	}
	user.name = "小王子"
	user.married = false
	fmt.Println(user)

	var p2 = new(animal)
	fmt.Println(p2)
	(*p2).name = "皮皮"
	(*p2).age = 2
	fmt.Printf("%v\n",p2)

	// 取结构体的地址进行实例化
	p3 := animal{}
	fmt.Printf("%T\n",p3)
	p3.name ="pi"
	fmt.Println(p3)

	// 结构体的初始化
	//1.键值对初始化
	p4:= person{
		name:"小王子",
		city:"武汉",
		age:20,
	}
	fmt.Printf("%#v\n",p4)

	p5 := &person{
		name: "yaya",
		city: "北京",
		age:  21,
	}
	fmt.Println(p5)
	//2.值的列表进行初始化
	p6:= person{
		"p6name","上海",20,info{"gril","752369@qq.com"},
	}
	fmt.Print(p6,"ppppp666")
	fmt.Printf("%#v\n",p6.info)
	fmt.Println(p6.info.sex)

	p10:= newPerson("腰子","武汉",27)
	fmt.Println(p10)
}

// 结构体的构造函数
func newPerson(name,city string,age int8) person{
	return person{
		name: name,
		city: city,
		age:  age,
	}
}
