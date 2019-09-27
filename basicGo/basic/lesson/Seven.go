package lesson

import (
	"encoding/xml"
	"fmt"
	"os"
)

// 人物档案
type person struct {
	Name string
	Age int
}

func Seven(){
	p := person{Name:"dav",Age:18}

	if data,err := xml.Marshal(p);err!=nil{
		fmt.Println(err)
		return
	}else {
		fmt.Println(string(data))
	}

	// 使用os.Args获取简单参数
	fmt.Println(os.Args)

}