package main

import (
	ginx "github.com/gin-gonic/gin"
	"github.com/vectorgo/mvc"
	"github.com/vectorgo/mvc/driver/gin"
	"strconv"
)

type userApi struct {}

func (u *userApi) list() string{
	return "hello nico"
}

var uapi = userApi{}

func main() {
	r := ginx.New()

	driver := gin.New(r)
	router := mvc.Router("/api/v1/users")
	router.Post("", driver.HandlerFunc(uapi.list))
	mvc.Use(router, driver).Complete()

	e := r.Run(":" + strconv.Itoa(8080))
	if e != nil{
		panic(e)
	}
}