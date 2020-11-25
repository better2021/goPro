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
	// 指定爬取起始、终止页
	var start,end int
	fmt.Println("请输入爬取的起始页(>=1)")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页(>=start)")
	fmt.Scan(&end)

	toWork(start,end)
}

func toWork(start,end int){
	fmt.Println("正在爬取%d到%d页",start,end)

	for i:=start;i<=end ;i++  {
		SpiderPage(i)
	}
}

func SaveFile(inx int,filmName,filmScore,peopleNum [][]string){
	fileURL,_ := os.Getwd()
	path := strings.Replace(fileURL, "\\", "/", -1)	// 替换\\为/
	fmt.Println(path,"--")
	dirName := path + "/" + "第" +  strconv.Itoa(inx) + "页.txt"
	f,err := os.Create(dirName)
	if err != nil{
		fmt.Println("create err",err)
		return
	}
	defer f.Close()

	n := len(filmName)
	fmt.Println(n,filmName[0][1],"---+---")
	fmt.Println(filmScore)
	fmt.Println(peopleNum)
	// 先获取电影名称
	f.WriteString("电影名称"+"\t\t\t"+"评分"+"\t\t"+"评分人数"+"\n")
	for i:=0;i<n ;i++{
		f.WriteString(filmName[i][1] + "\t\t\t" + filmScore[i][1] + "\t\t" + "\n")
	}
}


// 爬取单个页面数据
func SpiderPage(inx int)  {
	// 获取url
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((inx - 1) *25) + "&filter="

	result,err := HttpGet(url)
	if err != nil{
		fmt.Println("HttpGet err",err)
		return
	}

	//fmt.Println("result",result)

	// 解析、编译正则表达式 -- 名称
	ret:= regexp.MustCompile(`<img width="100" alt="(.+?)"`)
	// 提取需要信息
	filmName := ret.FindAllStringSubmatch(result,-1)

	// 解析、编译正则表达式 -- 分数
	pattern := `<span class="rating_num" property="v:average">(?s:(.*?))</span>`
	ret2:= regexp.MustCompile(pattern)

	// 提取需要信息
	filmScore := ret2.FindAllStringSubmatch(result,-1)

	// 解析、编译正则表达式 -- 人数
	ret3 := regexp.MustCompile(`<span width="100">(.*?)人评价</span>`)
	// 提取需要信息
	peopleNum := ret3.FindAllStringSubmatch(result,-1)

	SaveFile(inx,filmName,filmScore,peopleNum)
}

// 爬取指定url的页面，返回result
func HttpGet(url string) (result string,err error) {
	resp,err1 := http.Get(url)
	if err!=nil{
		err = err1
		return
	}
	defer resp.Body.Close()

	// 循环爬取整页的数据
	buf := make([]byte,4096)
	for  {
		n, err2 := resp.Body.Read(buf)
		if n == 0{
			break
		}

		if err2!=nil && err2 != io.EOF{
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}