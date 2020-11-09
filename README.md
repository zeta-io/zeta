# Introduce
A framework that helps other major Web frameworks work better.
# Feature
Zeta has the following features:
- less invasive
- use simple
- easily plug
- extensible
- practice is the sole criterion for testing truth
# Usage
The sample of Gin:
```go
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeta-io/ginx"
	"github.com/zeta-io/zeta"
)

func list(context context.Context, c1 *context.Context, args struct{
	Name string `json:"name" param:"query,name" validator:"required"`
}) (string, error){
	fmt.Println(args.Name)
	return "hello zeta", nil
}

func main() {
	router := zeta.Router("/api/:version/users")
	router.Get("", list)

	e := zeta.New(router, ginx.New(gin.New())).Run(":8080")
	if e != nil{
		panic(e)
	}
}
```
[More samples.](https://github.com/zeta-io/sample)
# Driver
- [Ginx](https://github.com/zeta-io/ginx): Gin driver for zeta. 