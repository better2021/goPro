package lesson

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func Three(){
	var r Retiever
	r = mock{"hello go"}
	r = Ret{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	fmt.Println(download(r))
}

// 定义接口
type Retiever interface {
	Get(url string) string
}

type mock struct {
	name string
}

// 定义Get方法
func (r mock) Get(url string) string {
	return  r.name
}

func download(r Retiever) string{
	return r.Get("http://www.baidu.com")  //r.Get调用接口中的方法
}

/*----------------------------------*/

type Ret struct {
	UserAgent string
	TimeOut time.Duration
}

func (r Ret) Get(url string) string{
	resp,err:=http.Get(url)
	if err!=nil{
		panic(err)
	}
	result,err := httputil.DumpResponse(resp,true)

	resp.Body.Close()

	if err != nil{
		panic(err)
	}
	return string(result)
}