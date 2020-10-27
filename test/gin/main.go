package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/first", func(c *gin.Context) {
		fmt.Println("first .........")
	})

	authorized := r.Group("/try")

	authorized.POST("/second", second)
	authorized.POST("/third", third)

	// 嵌套路由组
	testing := authorized.Group("testing")
	testing.GET("/forth", fourth)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func second(c *gin.Context) {
	//c.Request.RequestURI
	//

	fmt.Println("second .........")
}

func third(c *gin.Context) {
	fmt.Println("third .........")
}

func fourth(c *gin.Context) {
	fmt.Println("fourth .........")
}