package lesson

import (
	"fmt"
	"math"
)

func Five(){
	quyu()
}

// 取余
func quyu(){
	var a,b float64
	a = 19
	b = 5
	reminder := math.Mod(a,b)
	fmt.Println(reminder)
}