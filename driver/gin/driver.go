package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vectorgo/mvc"
	"github.com/vectorgo/mvc/http"
	"github.com/vectorgo/mvc/util/types"
	"reflect"
	"strings"
)

const ContextKey = "gin#context#key"

var(
	contextType = reflect.TypeOf(context.TODO())
	ginContextType = reflect.TypeOf(gin.Context{})
)

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

func process(ctx context.Context, c *gin.Context, call interface{}) {
	processor := newRequestParamsProcessor(c)

	typ := reflect.TypeOf(call)
	args := make([]reflect.Value, 0)
	for i := 0; i < typ.NumField(); i ++{
		f := typ.Field(i)
		ptr := false
		t := f.Type
		if t.Kind() == reflect.Ptr{
			ptr = true
			// handle as element type.
			t = t.Elem()
		}

		var target interface{}
		switch t {
		case contextType:
			target = ctx
		case ginContextType:
			target = *c
		default:
			name := f.Name
			source := ""
			if f.Tag.Get("param") == ""{
				ret, err := types.Convert("", t)
				if err != nil{
					panic(err)
				}
				target = ret
				break
			}
			params := strings.Split(f.Tag.Get("param"), ",")
			source = params[0]
			if len(params) > 1{
				name = params[1]
			}

			ret, err := processor.process(t, source, name)
			if err != nil{
				panic(err)
			}
			target = ret
		}
		if ptr{
			target = &target
		}
		args = append(args, reflect.ValueOf(target))
	}
	rets := reflect.ValueOf(call).Call(args)
	fmt.Println(rets)
}


