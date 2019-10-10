package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main()  {
	// 指定爬取起始、终止页
	var start,end int
	fmt.Println("请输入爬取的起始页(>=1)")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页(>=start)")
	fmt.Scan(&end)

	working(start,end)
}

// 爬取页面操作
func working(start,end int){
	fmt.Printf("正在爬取第%d页到%d页...\n",start,end)

	page := make(chan int)

	// 循环爬取每一页的数据
	for i:=start;i<=end;i++{
		go SpiderPage(i,page)
	}

	for i:=start;i<=end;i++{
		fmt.Println("第%s个页面爬取完成\n",<- page)
	}
}

func HttpGet(url string) (result string,err error){
	resp,err1 := http.Get(url)
	if err !=nil{
		err = err1
		return
	}
	defer resp.Body.Close()

	// 循环读取网页数据，传给调用者
	buf := make([]byte,4096)
	for{
		n,err2 := resp.Body.Read(buf)
		if n == 0{
			fmt.Println("读取网页完后")
			break
		}

		if err2 !=nil && err2!= io.EOF{
			err = err2
			return
		}
		// 累加每一次循环读取的buf数据，存入result
		result += string(buf[:n])
	}
	return
}

// 爬取单个页面的函数
func SpiderPage(i int,page chan int)  {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0" + strconv.Itoa((i-1)*50)
	result,err := HttpGet(url)
	if err !=nil{
		fmt.Println("HttpGet err",err)
		return
	}
	// fmt.Println("result=",result)
	// 将读取的整网页数据，保存成一个文件
	file,err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
	if err !=nil{
		fmt.Println("Create err",err)
		return
	}
	file.WriteString(result)
	file.Close() // 保存好一个文件就关闭一个文件

	// 将i写入通道page
	page <- i
}