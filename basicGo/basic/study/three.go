package study

import "fmt"

func Three()  {
	slice()
}

func slice(){
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,5,8,6}
	fmt.Println(arr1,arr2,arr3)

	for i := 0; i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for i,v := range arr2{
		fmt.Println(i,v)
	}
}