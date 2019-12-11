package main

import (
	"fmt"
	"net/http"
)

// 路由
func Run(){
	http.HandleFunc("/",IndexView)
	http.HandleFunc("/upload",UploadView)
	http.HandleFunc("/api/upload",ApiUpload)
	fmt.Println("run at 8080")
	http.ListenAndServe(":8088",nil)
}