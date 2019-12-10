package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func IndexView(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello"))
}

// 上传
func Upload(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method)
	body,_ := ioutil.ReadFile("./upload.html")
	// fmt.Println(string(body),"--")
	// GET
	if r.Method == "GET"{
		w.Write([]byte(string(body)))
		return
	}
	// POST
	if r.Method =="POST"{
		file,header,err:=r.FormFile("file")
		if err != nil{
			// 有错误
			w.Write([]byte("文件上传有误："+err.Error()))
			return
		}
		os.Mkdir("./images",0666)
		out,err := os.Create("./images/" + header.Filename)
		if err !=nil{
			io.WriteString(w,"文件创建失败:"+err.Error())
		}
		_,err = io.Copy(out,file)
		if err !=nil{
			io.WriteString(w,"文件保存失败:"+err.Error())
		}
	}
}

// 返回指定的图片
func ImageView(w http.ResponseWriter,r *http.Request){
	r.ParseForm() // 把url 或者 form 表单的数据解析到对应的容器 r.From
	name := r.Form.Get("name")
	fmt.Printf("name:%s",name)
}

func main()  {
	fmt.Println("hello")
	http.HandleFunc("/upload",Upload)
	http.HandleFunc("/index",IndexView)
	http.HandleFunc("/image",ImageView)
	fmt.Println("run at 8081")
	http.ListenAndServe(":8081",nil)
}
