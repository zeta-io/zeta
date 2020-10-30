package main

import (
	"fmt"
	"net/url"
)

func main(){
	query := "params=fo,fu,n%2Cico"
	values, err := url.ParseQuery(query)
	if err != nil{
		panic(err)
	}
	for _, vs := range values{
		for _, v := range vs{
			fmt.Println(v)
		}
	}
}
