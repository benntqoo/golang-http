package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {
	//註冊路由&回傳方法
	http.HandleFunc("/", helloWorld)

	http.HandleFunc("/hi",hi)

	//監聽位置 假設監聽本機位置直接 :port 號
	//指定 ip 位置 ex. 162.168.10.2:8080
	//第二參數為 serverMux nil == defaultServerMux
	http.ListenAndServe(":8080", nil)
}

//路由的回傳方法
func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func hi(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hi golang web app")
}