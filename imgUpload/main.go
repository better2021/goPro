package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func IndexView(w http.ResponseWriter,r *http.Request){
	w.Write([]byte(`<p><a href="/list">列表</a> <a href="/upload">上传</a></p>`))
}

// 上传
func Upload(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method)

	// GET
	if r.Method == "GET"{
		body,_ := ioutil.ReadFile("./upload.html")
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
		t := header.Header.Get("Content-Type")
		fmt.Println(t,"--")
		if !strings.Contains(t,"image"){
			io.WriteString(w,"请上传图片")
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
			return
		}
		//io.WriteString(w,"文件保存成功")
		http.Redirect(w,r,"/image?name="+header.Filename,302) //302临时重定向
	}
}

// 返回指定的图片
func ImageView(w http.ResponseWriter,r *http.Request){
	r.ParseForm() // 把url 或者 form 表单的数据解析到对应的容器 r.From
	name := r.Form.Get("name")
	fmt.Printf("name:%s",name)
	f,err:= os.Open("./images/" + name)  // 预览地址 http://localhost:8081/image?name=photo_2019-12-07_15-18-37.jpg
	if err != nil{
		w.WriteHeader(404)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Type","image")
	io.Copy(w,f)
}

// 详细
func DetailView(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	name := r.Form.Get("name")
	fmt.Println(name)
	html := loadHtml("./detail.html")
	html = bytes.Replace(html,[]byte("@src"),[]byte("/image?name="+name),1)
	w.Write(html)

	//io.WriteString(w,"图片的名称是："+ name)
}

// 加载html
func loadHtml(name string) []byte{
	f,err := os.Open(name)
	if err != nil{
		return []byte("error")
	}
	defer f.Close()
	buf,err := ioutil.ReadAll(f)
	if err !=nil{
		return []byte("errors")
	}
	return buf
}

func ListView(w http.ResponseWriter,r *http.Request){
	html := loadHtml("./list.html")
	names,err := ioutil.ReadDir("./images")
	if err!=nil{
		w.Write([]byte("erros"))
		return
	}
	temp := ""
	for i:=0;i<len(names);i++{
		log.Println(names[i].Name(),err)

		temp += `<li><a href="/detail?name=`+ names[i].Name() +`"><img src="/image?name=`+ names[i].Name() +`" alt="图片"></a></li>`
	}
	html = bytes.Replace(html,[]byte("@html"),[]byte(temp),1)
	w.Write(html)
}

func main()  {
	fmt.Println("hello")
	http.HandleFunc("/",IndexView)
	http.HandleFunc("/upload",Upload)
	http.HandleFunc("/image",ImageView)
	http.HandleFunc("/list",ListView)
	http.HandleFunc("/detail",DetailView)
	fmt.Println("run at 8081")
	http.ListenAndServe(":8081",nil)
}
