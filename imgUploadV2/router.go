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
	http.Handle("/static/",http.FileServer(http.Dir("static/")))
	fmt.Println("run at 8080")
	http.ListenAndServe(":8088",nil)
}