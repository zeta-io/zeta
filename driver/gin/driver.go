package gin

import (
	"context"
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
		call := m
		handleFunc = append(handleFunc, func(c *gin.Context){
			ctx := context.WithValue(context.Background(), ContextKey, c)
			if call == nil{
				panic("handler func args is nil.")
			}
			if reflect.TypeOf(call).Kind() != reflect.Func{
				panic("handler func type must be func.")
			}
			if c.IsAborted(){
				return
			}
			rets, err := process(ctx, c, call)
			if err != nil{
				d.r(c, err.Error(), err)
				return
			}
			if len(rets) > 0{
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
		})
	}
	d.e.Handle(string(method), url, handleFunc...)
}

func process(ctx context.Context, c *gin.Context, call interface{}) ([]reflect.Value, error){
	processor, err := newRequestParamsProcessor(c)
	if err != nil{
		return nil, err
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
			target, err = processRequestParams(processor, in, ptr)
			if err != nil{
				return nil, err
			}
		}
		args = append(args, target)
	}
	return reflect.ValueOf(call).Call(args), nil
}

func processRequestParams(processor *requestParamsProcessor, in reflect.Type, ptr bool) (reflect.Value, error){
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

		ret, ok, err := processor.process(ft, source, name)
		if err != nil{
			return obj, err
		}
		if (!ok || ret == nil) && ptr{
			target := reflect.New(f.Type).Elem()
			obj.FieldByName(f.Name).Set(target)
		}else{
			target := reflect.New(ft).Elem()
			if ret != nil{
				target.Set(reflect.ValueOf(ret))
			}
			if ptr{
				target = target.Addr()
			}
			obj.FieldByName(f.Name).Set(target)
		}
	}
	if ptr{
		obj = obj.Addr()
	}
	return obj, nil
}


