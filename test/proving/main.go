package main

import "github.com/vectorgo/mvc"

type UserApi struct {
	getUser int `api:"/api/v1/users" method:"post" `
}


func get(){

}

func handleFunc(uri string, fn interface{}){

}

func main() {
	//test := Test{}
	//t := reflect.TypeOf(test)
	//methodNum := t.NumMethod()
	//
	//for i := 0; i < methodNum; i ++{
	//	t.Method(i).
	//}
	//t.Field(1).Tag
	//handleFunc("abc", get)

	mvc.New(mvc.Config{}).Router().Post("", nil).Group("u").Post("a", nil)
}