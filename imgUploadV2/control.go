package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// 控制器

// 首页
func IndexView(w http.ResponseWriter,r *http.Request){
	html := LoadHtml("./views/index.html")
	w.Write(html)
}

// 上传页面
func UploadView(w http.ResponseWriter,r *http.Request){
	html := LoadHtml("./views/upload.html")
	w.Write(html)
}

// 图片上传
func ApiUpload(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	f,h,err := r.FormFile("file")
	if err !=nil{
		io.WriteString(w,"上传错误")
		return
	}
	t:= h.Header.Get("Content-Type")
	if !strings.Contains(t,"image"){
		io.WriteString(w,"文件类型错误")
		return
	}

	os.Mkdir("./static",0666)
	name := time.Now().Format("2006-01-02150405") +  path.Ext(h.Filename) // 获取文件后缀
	fmt.Println(name,"--")
	out,err := os.Create("./static/" + name)
	// fmt.Println(out,err,"++")
	if err != nil{
		io.WriteString(w,"文件创建错误")
		return
	}
	io.Copy(out,f)
	out.Close()
	f.Close()
	mod := Info{
		Name:h.Filename,
		Path:name,
		Note:r.Form.Get("note"),
		CreateTime:time.Now().Unix(),
	}
	InfoAdd(&mod)
}

// 详情页面
func DetailView(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	url:=r.Form.Get("url")
	html:= LoadHtml("./view/detail.html")
	bytes.Replace(html,[]byte("@src"),[]byte(url),1)
	w.Write(html)
}

// 加载Html文件
func LoadHtml(name string) []byte{
	buf,err := ioutil.ReadFile(name)
	if err !=nil{
		return []byte("error")
	}
	return buf
}