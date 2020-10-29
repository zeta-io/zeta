package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vectorgo/mvc"
	"github.com/vectorgo/mvc/http"
	"reflect"
)

const ContextKey = "gin#context#key"

type Driver struct {
	e *gin.Engine
}

func New(e *gin.Engine) Driver{
	return Driver{e}
}

func (d Driver) Option(m *mvc.Mvc){
	m.Driver(d)
}

func (d Driver) Handle(method http.Method, url string, middleware ...mvc.HandlerFunc){
	handleFunc := make([]gin.HandlerFunc, 0)
	for _, m := range middleware{
		handleFunc = append(handleFunc, func(c *gin.Context){
			m(context.WithValue(context.Background(), ContextKey, c))
		})
	}
	d.e.Handle(string(method), url, handleFunc...)
}

func (d Driver) HandlerFunc(call interface{}) mvc.HandlerFunc{
	if call == nil{
		panic("handler func args is nil.")
	}
	if reflect.TypeOf(call).Kind() != reflect.Func{
		panic("handler func type must be func.")
	}
	return func(c context.Context) {
		o := c.Value(ContextKey)
		if o == nil{
			panic("gin context is nil.")
		}
		gc, ok := o.(*gin.Context)
		if ! ok{
			panic(fmt.Sprintf("can't cast %v to *gin.Context.", o))
		}
		process(c, gc, call)
	}
}

func process(c context.Context, gc *gin.Context, call interface{}) {
	ct := reflect.TypeOf(call)
	for i := 0; i < ct.NumField(); i ++{
		f := ct.Field(i)

	}
}

