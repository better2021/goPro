package lesson

import (
	"fmt"
	"strconv"
	"strings"
)

func Six(){
	// 是否包含
	s := "hello world"
	fmt.Println(strings.Contains(s,"hello"),strings.Contains(s,"?"))

	// 索引，base 0
	fmt.Println(strings.Index(s,"o"))

	ss := "1#2#345"
	splitStr := strings.Split(ss,"#") // 用#号分割字符串
	fmt.Println(splitStr)
	fmt.Println(strings.Join(splitStr,"#")) // 合并字符串

	fmt.Println(strings.HasPrefix(s,"he")) // s中是否含有he的前缀

	fmt.Println(strings.HasSuffix(s,"ld")) // s中是否含有ld的后缀

	/*字符串转换*/
	fmt.Println(strconv.Itoa(10)) // 转化为整形
	fmt.Println(strconv.Atoi("789")) // 转化为字符串

	fmt.Println(strconv.ParseBool("false")) // 转换bool类型
	fmt.Println(strconv.ParseFloat("3.14",64)) // 转化为浮点型
}


