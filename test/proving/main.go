package main

import (
	"context"
	"fmt"
	ginx "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zeta-io/zeta"
	"github.com/zeta-io/zeta/driver/gin"
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
	B *bool `param:"query,b"`
	Version string `param:"path,version"`
	V1 *string `param:"path,version"`
	Size int  `param:"query,size" validate:"gte=0,lte=20"`
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
	fmt.Println("v1 is null ? ", args1.V1 == nil)
	fmt.Println("version is ", args1.Version)
	fmt.Println("size is", args1.Size)
	return "hello nico", nil
}

var uapi = userApi{}

func main() {
	router := zeta.Router("/api/:version/users")
	router.Post("", uapi.list)

	e := zeta.New(router, gin.New(ginx.New()).Response(func(c *ginx.Context, data interface{}, err error){
		if err != nil{
			if validateErrs, ok := err.(validator.ValidationErrors); ok && len(validateErrs) > 0{
				validateErr := validateErrs[0]
				c.JSON(200, fmt.Sprintf("%s param err, reason: %v", validateErr.Namespace(), validateErr))
				c.Abort()
			}
			return
		}
		c.JSON(200, data)
		c.Abort()
	})).Run(":8080")
	if e != nil{
		panic(e)
	}
}