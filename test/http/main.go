package http

import (
	"log"
	http1 "net/http"
)

func sayhelloGolang(w http1.ResponseWriter, r *http1.Request) {

}

func main() {
	http1.HandleFunc("/", sayhelloGolang) //设置访问的路由
	err := http1.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	http1.DetectContentType()
}
