package lesson

import "fmt"

func One(){
	m1:= map[int]string{1:"a",2:"b",3:"c",4:"d"}
	fmt.Println(m1,m1[2])

	m2 := make(map[string]int)
	for k,v :=range m1{
		m2[v] = k
	}
	m2["a"] = 18
	fmt.Println(m2,m2["a"])

	addRe := add(2,5,3)
	fmt.Println(addRe)

	param(2,6,3)

	s1:=[]int{1,2,3}
	aa(s1)

}

func add(a,b,c int) int{
	return  a + b + c
}

func param(x ...int){
	fmt.Println(x)
	fmt.Println(len(x))
}

func aa(s []int){
	s[0] = 5
	s[2] = 6
	fmt.Println(s,"--")
}