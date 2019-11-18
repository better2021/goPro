package study

import (
	"fmt"
)

func One(){
	var a,b int = 10 ,20
	var s string

	s = "coco"
	fmt.Printf("%d %d %q\n",a,b,s)

	e,f,j := 3,"qq",true
	fmt.Printf("%d %s %v\n",e,f,j)

	var (
		aa = 3
		bb = "qwert"
		cc = false
	)
	fmt.Printf("%d %s %t\n",aa,bb,cc)

	const (
		x = 10
		y = 15
	)
	fmt.Print(x,y)
}