package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

func httpParams(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认不解析
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("schema:", r.URL.Scheme)
	fmt.Println(r.Form["name"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v,","))
	}
	fmt.Fprintf(w, "Hello HH") // 输出到客户端的
}

func startHttp() {
	http.HandleFunc("/", httpParams) // 设置访问的路由
	err := http.ListenAndServe(":9990", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
