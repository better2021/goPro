package lesson

import (
	"fmt"
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
}