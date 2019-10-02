package lesson

import (
	"fmt"
)

// 人物档案
type person struct {
	Name string
	Age int
}

func Seven(){
	//p := person{Name:"dav",Age:18}
	//
	//if data,err := xml.Marshal(p);err!=nil{
	//	fmt.Println(err)
	//	return
	//}else {
	//	fmt.Println(string(data))
	//}
	//
	//// 使用os.Args获取简单参数
	//fmt.Println(os.Args)

	// Range 遍历
	nums := []int{2,3,4,5,6}
	sum:=0
	for _,num := range nums{
		sum += num
	}
	fmt.Println("sum:",sum)

	for i,num := range nums{
		if num == 3{
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s->%s\n",k,v)
	}

}