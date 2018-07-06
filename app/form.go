package main

import (
	"net/http"
	"fmt"
	"log"
)

func putOut(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	fmt.Println(r.Form)
	fmt.Println("method: ", r.Method)
	var returnStr string
	if r.Method == "POST" {
		returnStr = r.Form["username"][0]
	}
	fmt.Fprintf(w, returnStr)
}

func formHttp() {
	http.HandleFunc("/myusername", putOut)
	err := http.ListenAndServe(":9009", nil)
	if err != nil {
		log.Fatal(err)
	}
}
