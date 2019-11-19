package study

import (
	"fmt"
	"unicode/utf8"
)

func Three()  {
	// slice()

	// appe()
	as()
}

func slice(){
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,5,8,6}
	fmt.Println(arr1,arr2,arr3)

	fmt.Println(arr3[1:3])
	fmt.Println(arr3[:2])
	fmt.Println(arr3[2:])

	for i := 0; i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for i,v := range arr2{
		fmt.Println(i,v)
	}
}

func appe()  {
	var s []int

	for i:=0;i<50 ;i++{
		s = append(s,2*i+1)
	}
	fmt.Println(s,len(s))
}

func as(){
	s := "Yes慕课网"
	 for _,b := range []byte(s){
	 	fmt.Printf("%x\n",b)
	 }
	 fmt.Println(s)

	 for i,ch := range s{
	 	fmt.Printf("%d %X",i,ch)
	 }

	 bytes := []byte(s)
	 for len(bytes) > 0{
	 	ch,size := utf8.DecodeRune(bytes)
	 	bytes = bytes[size:]
	 	fmt.Printf("%c\n",ch)
	 }
}