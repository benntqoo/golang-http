package main

import (
	"net/http"
	"io"
	"fmt"
	"time"
)

type webHandler struct {
}

func (h *webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler,ok:= muxMap[r.URL.String()];ok{
		handler(w,r)
		return
	}

	fmt.Fprintf(w, "ServerMux Handler  URL = %s ",r.URL.String())
}

var muxMap map[string]func(w http.ResponseWriter,r *http.Request)

func main() {
	//註冊路由&回傳方法
	muxMap = make( map[string]func(w http.ResponseWriter,r *http.Request))
	muxMap["/hello"]=helloWorld
	muxMap["/hi"]=hi


	// create http Server struct
	server:=http.Server{
		Addr:":8080",//ip addr
		Handler:&webHandler{},//server mux
		ReadHeaderTimeout:10*time.Second,//set timeout
	}

	//監聽位置
	server.ListenAndServe()
}

//路由的回傳方法
func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi golang web app")
}
