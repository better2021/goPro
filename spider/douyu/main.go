package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main()  {
	url := "https://www.douyu.com/g_yz"
	// 爬取整个页面，将整个页面全部保存在result
	result,err := HttpGet(url)
	if err !=nil{
		fmt.Println("HttpGet err",err)
		return
	}
	//fmt.Println(result)
	// 解析编译正则
	ret := regexp.MustCompile(`src="(?s:(.*?))"`)
	// 提取每一张图片的url
	alls := ret.FindAllStringSubmatch(result,50)  // 获取30张图片地址,-1表示全部

	page := make(chan int)
	n := len(alls)

	for idx,imgUrl := range alls{
		fmt.Println("imgUrl",imgUrl[1])
		go SaveImg(idx,imgUrl[1],page)
	}

	for i:=0;i<n;i++{
		fmt.Println("下载第%d张图片完后\n",<-page)
	}

	// 防止主go程序退出
	//for {
	//	runtime.GC()
	//}
}


// 获取一个网页的所有内容，result返回
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

// 保存图片
func SaveImg(idx int,url string,page chan int){
	fileURL,_ := os.Getwd()
	path := strings.Replace(fileURL, "\\", "/", -1)	// 替换\\为/
	//fmt.Println(path,"--")
	dirName := path + "/img/" +  strconv.Itoa(idx + 1) + ".jpg"
	f,err := os.Create(dirName)
	if err !=nil{
		fmt.Println("HttpGet err",err)
		return
	}
	defer f.Close()

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
			break
		}

		if err2 !=nil && err2!= io.EOF{
			err = err2
			return
		}

		f.Write(buf[:n])
	}
	page <-idx
}