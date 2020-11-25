package study

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"runtime"
)

func Web(){
	http.HandleFunc("/hello",helloworld)

	files:= http.FileServer(http.Dir("./basic"))
	fmt.Println(files)
	http.Handle("/static",http.StripPrefix("/static/",files))

	// 访问隐私页面
	http.HandleFunc("/profile", profile)

	// 设置session
	http.HandleFunc("/setSession", setSession)

	server := http.Server{
		Addr:              ":8080",
	}
	server.ListenAndServe()
}

func helloworld(w http.ResponseWriter,r *http.Request)  {
	// 默认不会解析，需要先解析表单
	err := r.ParseForm()
	if err != nil{
		fmt.Println("参数解析出错",err)
		return
	}

	type Person struct {
		Name string
		Age int
	}

	p:= Person{
		Name: "coco",
		Age:  50,
	}

	data,_:=json.Marshal(&p) // JSON序列化与反序列化需要使用encoding/json包
	fmt.Println(string(data))

	fmt.Println("path",r.URL.Path)
	fmt.Println(r.Form)
	fmt.Fprint(w,"hello")

	fmt.Println("cpu核数",runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Gosched)

	str := `string我爱学习`
	fmt.Println(str,len(str))
	// rune 等同于int32,常用来处理unicode或utf-8字符
	rr := []rune(str)
	fmt.Println(rr)
	for i:=0;i<len(rr);i++{
		fmt.Printf("%c \n",rr[i])
	}
}

// 利用cookie方式创建session，秘钥为 mykey
var store = sessions.NewCookieStore([]byte("mykey"))

func setSession(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "sid")
	session.Values["username"] = "张三"
	session.Save(r, w)
}

func profile(w http.ResponseWriter, r *http.Request){

	session, _ := store.Get(r, "sid")

	if session.Values["username"] == nil {
		fmt.Fprintf(w, `未登录，请前往 localhost:8080/setSession`)
		return
	}

	fmt.Fprintf(w, `已登录，用户是：%s`, session.Values["username"])
	return
}