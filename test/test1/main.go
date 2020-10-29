package main

import (
	"context"
	"fmt"
	"reflect"
)

func test(f interface{}){
	f_type := reflect.TypeOf(f)
	fmt.Println(f_type)
	fmt.Println(f_type.Kind())
	f_value := reflect.ValueOf(f)
	fmt.Println(f_value)
	f_value.Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(struct {

	}{})})
}

func userList(c context.Context, args struct{}){
	fmt.Println(c, args)
}

func main() {
	test(userList)
}
