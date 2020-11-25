package main

import (
	"fmt"
	"net/http"
)

// 路由
func Run(){
	http.HandleFunc("/",IndexView)
	http.HandleFunc("/upload",UploadView)
	http.HandleFunc("/list",ListView)
	http.HandleFunc("/api/upload",ApiUpload)
	http.HandleFunc("/api/list",ApiList)
	http.HandleFunc("/detail",DetailView)
	http.HandleFunc("/api/delete",ApiDelete)
	/*
	http.FileServer 返回的Handler将会进行查找，并将与文件夹或文件系统有关的内容以参数的形式返回给你（在这里你将"static"作为静态文件的根目录）。
	因为你的"example.txt"文件在静态目录中，你必须定义一个相对路径去获得正确的文件路径。
	*/
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	fmt.Println("run at 8088")
	http.ListenAndServe(":8088",nil)
}