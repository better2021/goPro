package lesson

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func Time()  {
	d := time.Now()
	fmt.Println(d)

	local,err := time.LoadLocation("Asia/Shanghai")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(local)

	year, month, day := d.Date()
	fmt.Println(year,month,day)

	hour, minute, second := d.Clock()
	fmt.Println(hour, minute, second)

	fmt.Println(d.Year(),d.Month())
	fmt.Println(d.Weekday())
	fmt.Println(d.Unix())
	fmt.Println(d.Local())
	fmt.Println(d.Location())

	fmt.Println(strconv.Atoi("123")) // 字符串转整形 string => int
	fmt.Println(strconv.Itoa(123)) // 整形转字符串


	 // 以xx结尾的字节,返回true或false
	fmt.Println(bytes.HasPrefix([]byte("Gopher"),[]byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"),[]byte("er")))

	// 以xx结尾的字节,返回true或false
	fmt.Println(bytes.HasSuffix([]byte("Gopher"),[]byte("Go")))
	fmt.Println(bytes.HasSuffix([]byte("Gopher"),[]byte("er")))

	fmt.Println(len("12as"))
}