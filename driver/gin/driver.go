package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vectorgo/mvc"
	"github.com/vectorgo/mvc/http"
	"reflect"
	"strings"
)

const ContextKey = "gin#context#key"

var(
	contextType = "context.Context"
	ginContextType = "gin.Context"
)

type Driver struct {
	e *gin.Engine
	r func(c *gin.Context, data interface{}, err error)
}

func New(e *gin.Engine, r func(c *gin.Context, data interface{}, err error)) Driver{
	return Driver{e: e, r: r}
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
	return func(ctx context.Context) {
		o := ctx.Value(ContextKey)
		if o == nil{
			panic("gin context is nil.")
		}
		c, ok := o.(*gin.Context)
		if ! ok{
			panic(fmt.Sprintf("can't cast %v to *gin.Context.", o))
		}
		rets := process(ctx, c, call)
		var data interface{}
		var err error
		for _, ret := range rets{
			if e, ok := ret.Interface().(error); ok{
				err = e
			}else if data == nil{
				data = ret.Interface()
			}
		}
		d.r(c, data, err)
	}
}

func process(ctx context.Context, c *gin.Context, call interface{}) []reflect.Value{
	processor, err := newRequestParamsProcessor(c)
	if err != nil{
		panic(err)
	}
	typ := reflect.TypeOf(call)
	args := make([]reflect.Value, 0)
	for i := 0; i < typ.NumIn(); i ++{
		in := typ.In(i)
		ptr := false
		if in.Kind() == reflect.Ptr{
			ptr = true
			// handle as element type.
			in = in.Elem()
		}

		var target reflect.Value
		switch in.String() {
		case contextType:
			target = reflect.ValueOf(ctx)
			if ptr{
				target = reflect.ValueOf(&ctx)
			}
		case ginContextType:
			target = reflect.ValueOf(*c)
			if ptr{
				target = reflect.ValueOf(c)
			}
		default:
			if in.Kind() != reflect.Struct{
				continue
			}
			target = processRequestParams(processor, in, ptr)
		}
		args = append(args, target)
	}
	return reflect.ValueOf(call).Call(args)
}

func processRequestParams(processor *requestParamsProcessor, in reflect.Type, ptr bool) reflect.Value{
	obj := reflect.New(in).Elem()
	for i := 0; i < in.NumField(); i ++{
		f := in.Field(i)
		name := f.Name
		source := ""
		if f.Tag.Get("param") == ""{
			continue
		}
		params := strings.Split(f.Tag.Get("param"), ",")
		source = params[0]
		if len(params) > 1{
			name = params[1]
		}

		ft := f.Type
		ptr := false
		if ft.Kind() == reflect.Ptr{
			ptr = true
			ft = ft.Elem()
		}

		ret, err := processor.process(ft, source, name)
		if err != nil{
			panic(err)
		}
		target := reflect.New(ft).Elem()
		target.Set(reflect.ValueOf(ret))
		if ptr{
			target = target.Addr()
		}
		obj.FieldByName(f.Name).Set(target)
	}
	if ptr{
		obj = obj.Addr()
	}
	return obj
}


