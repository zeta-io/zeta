package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestDriver_HandlerFunc(t *testing.T) {
	driver := Driver{}
	handlerFunc := driver.HandlerFunc(func(c context.Context, gc *gin.Context, name string, age int, args struct{
		Title string
	}) string{
		return "hello nico"
	})
	handlerFunc(context.WithValue(context.Background(), ContextKey, &gin.Context{}))
}