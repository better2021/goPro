package main

import (
	"fmt"
	"reflect"
)

type Cat struct {

}

type Dog struct {
	Name string
}

type student struct {
	Name string `json:"name"`
	Score int `json:"score"`
}

/*
通过反射可以获取到数据的类型
interface{} 空接口可以表示任何的类型
*/
func reflectType(x interface{}){
	// 1.通过类型断言
	// 2.借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
}

func reflectValue(x interface{})  {
	v:=reflect.ValueOf(x)
	fmt.Println(v)
	k:=v.Kind() // 拿到值对应的类型种类
	// fmt.Println(k)

	switch k {
		case reflect.Float32:
			// 把反射取到的值转换为一个int32类型的变量
			ret:=float32(v.Float())
			fmt.Println(ret)
		case reflect.Int32:
			ret:=int32(v.Int())
			fmt.Println(ret)
	}
}

func main()  {
	//var a float64 = 3.14
	//reflectType(a)
	//var b int8 = 10
	//reflectType(b)
	//
	//var c Cat
	//reflectType(c)
	//var d Dog
	//reflectType(d.Name)
	//
	//var e []int
	//var f []string
	//reflectType(e)
	//reflectType(f)

	//var aa int32 = 16
	//var bb string = "yaya"
	//reflectValue(aa)
	//reflectValue(bb)

	stu1:= student{
		Name:  "小王子",
		Score: 80,
	}

	// 通过反射去获取结构体重所以字段信息
	t:=reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n",t.Name(),t.Kind())

	for i:=0;i<t.NumField();i++{
		// i为结构体字段的索引
		fileObj := t.Field(i)
		fmt.Println(fileObj.Name,fileObj.Type,fileObj.Tag)
		fmt.Println(fileObj.Tag.Get("json"),"---")
	}

	//  根据名字去取结构体中的字段
	filed,ok:=t.FieldByName("Score")
	if ok{
		fmt.Println(filed.Name,filed.Type)
	}
}