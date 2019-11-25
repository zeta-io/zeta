package main

import (
	"fmt"
	"github.com/vectorgo/mvc/core/json"
	"github.com/vectorgo/mvc/core/parser/http"
	"log"
	http1 "net/http"
)

func sayhelloGolang(w http1.ResponseWriter, r *http1.Request) {
	parser := http.Parser{}
	requestInfo, err:= parser.Parse(r)
	fmt.Println(err)
	fmt.Println(json.ToJsonIgnoreError(requestInfo))
}

func main() {
	http1.HandleFunc("/", sayhelloGolang) //设置访问的路由
	err := http1.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
