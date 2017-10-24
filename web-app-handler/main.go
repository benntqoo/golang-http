package main

import (
	"net/http"
	"io"
	"fmt"
)

type webHandler struct {
}

func (h *webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ServerMux Handler")
}

func main() {
	//自訂義 serverHandler
	mux := http.NewServeMux()

	//註冊路由&回傳方法
	mux.Handle("/", &webHandler{})

	mux.HandleFunc("/hello", helloWorld)
	mux.HandleFunc("/hi", hi)

	//監聽位置
	http.ListenAndServe(":8080", mux)
}

//路由的回傳方法
func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi golang web app")
}
