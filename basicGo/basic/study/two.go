package study

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Two(){
	fileURL,_ := os.Getwd()
	fmt.Print(fileURL,"--")

	path := strings.Replace(fileURL, "\\", "/", -1)	// 替换\\为/
	filename := path + "/study/abc.txt"
	if contents,err := ioutil.ReadFile(filename);err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}

	// fmt.Print(grade(10),grade(85),grade(93),grade(102))

	fmt.Print(converToBin(5))

	printFile(filename)

	fmt.Println(eval(5,3,"+"))
	fmt.Println(eval(3,6,"*"))
	fmt.Println(eval(5,3,"-"))
	fmt.Println(eval(15,3,"/"))

	fmt.Println(sum(1,9))
}

/*
	switch
*/
func grade(score int) string{
	g := ""
	switch  {
	case score < 60:
		g =  "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <100:
		g = "A"
	default:
		fmt.Printf("Wrong score: %d", score)
		panic("wrong")
	}
	return g
}

/*
for
*/
func converToBin(n int) string{
	result := ""
	for ; n > 0 ; n/=2{
		lsb := n%2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	fmt.Println(scanner,"--**")
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}

/*
 函数
*/
func eval(a,b int,op string) int {
	result := 0
	switch op {
		case "+":
			result =  a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result =  a / b
		default:
			fmt.Println("err")
	}
 	return result
}

/*
可变参数
*/
func sum(numbers ...int) int{
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}
