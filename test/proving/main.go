package main

import (
	"context"
	"fmt"
	ginx "github.com/gin-gonic/gin"
	"github.com/vectorgo/mvc"
	"github.com/vectorgo/mvc/driver/gin"
)

type userApi struct {}

func (u *userApi) list(context context.Context, c1 *context.Context, c *ginx.Context, c2 ginx.Context, args *struct{
	Name string `json:"name" param:"query,name"`
	Age *int `json:"age" param:"query,age"`
}, args1 struct{
	Name string `json:"name" param:"query,name"`
	Age *int `json:"age" param:"query,age"`
	NA  *string `json:"na" param:"query,na"`
	NB  *string `json:"nb" param:"query,nb"`
}) (string, error){
	fmt.Println(context)
	fmt.Println(c1)
	fmt.Println(c)
	fmt.Println(c2)
	fmt.Println(args.Name)
	fmt.Println(args.Age == nil)
	fmt.Println(args1.Name)
	fmt.Println(args1.Age == nil)
	fmt.Println(args1.NA == nil)
	fmt.Println(args1.NA)
	fmt.Println(args1.NB == nil)
	fmt.Println(args1.NB)
	return "hello nico", nil
}

var uapi = userApi{}

func main() {
	driver := gin.New(ginx.New(), mvc.DefaultSerial(), func(c *ginx.Context, data interface{}, err error){
		c.JSON(200, data)
		c.Abort()
	})
	router := mvc.Router("/api/v1/users")
	router.Post("", uapi.list)

	e := mvc.New(router, driver).Run(":8080")
	if e != nil{
		panic(e)
	}
}