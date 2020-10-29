package mvc

import "github.com/vectorgo/mvc/http"

type group struct {
	r     *router
	url        string
	middleware  []HandlerFunc
}

func (g *group) Use(middleware ...HandlerFunc) *group{
	g.middleware = append(g.middleware, middleware...)
	return g
}

func (g *group) Handle(method http.Method, url string, middleware ...HandlerFunc) *router {
	return g.r.Handle(method, g.url + url, append(g.middleware, middleware...)...)
}

func (g *group) Get(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodGet, url, middleware...)
}

func (g *group) Post(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodPost, url, middleware...)
}

func (g *group) Put(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodPut, url, middleware...)
}

func (g *group) Delete(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodDelete, url, middleware...)
}

func (g *group) Patch(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodPatch, url, middleware...)
}

func (g *group) Head(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodHead, url, middleware...)
}

func (g *group) Options(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodOptions, url, middleware...)
}

func (g *group) Connect(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodConnect, url, middleware...)
}

func (g *group) Trace(url string, middleware ...HandlerFunc) *router {
	return g.Handle(http.MethodTrace, url, middleware...)
}
