// Router
package mvc

import "github.com/vectorgo/mvc/http"

type RouterType uint8

const (
	RouterTypeGroup  RouterType = 1
	RouterTypeRouter
)

var Router = router{}

type Middleware func(c Context)

type router struct {
	middleware  []Middleware
	mappings    []mapping
	groups      []group
}

type mapping struct {
	url        string
	call       interface{}
	method     http.Method
	middleware  []Middleware
}

type group struct {
	url        string
	middleware  []Middleware
	mappings    []mapping
}

func (r *router) Group(url string, middleware ...Middleware) *group {
	g := &group{
		middleware: middleware,
		url: url,
	}
	if r.groups == nil{
		r.groups = make([]group, 0)
	}
	r.groups = append(r.groups, g)
	return g
}

func (r *router) Handle(method http.Method, url string, call interface{}, middleware ...Middleware) *router {
	if r.mappings == nil {
		r.mappings = make([]mapping, 0)
	}
	r.mappings = append(r.mappings, mapping{
		method:     method,
		url:        url,
		call:       call,
		middleware: middleware,
	})
	return r
}

func (r *router) Get(url string, call interface{}, middleware ...Middleware) *router {
	return r.Handle(http.MethodGet, url, call, middleware...)
}

func (r *router) Post(url string, call interface{}, middleware ...Middleware) *router {
	return r.Handle(http.MethodPost, url, call, middleware...)
}

func (r *router) Put(url string, call interface{}, middleware ...Middleware) *router {
	return r.Handle(http.MethodPut, url, call, middleware...)
}

func (r *router) Delete(url string, call interface{}, middleware ...Middleware) *router {
	return r.Handle(http.MethodDelete, url, call, middleware...)
}
