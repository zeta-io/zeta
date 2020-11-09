package zeta

type group struct {
	r     *router
	url        string
	middleware  []HandlerFunc
}

func (g *group) Use(middleware ...HandlerFunc) *group{
	g.middleware = append(g.middleware, middleware...)
	return g
}

func (g *group) Group(url string, middleware ...HandlerFunc) *group {
	return &group{
		middleware: append(g.middleware, middleware...),
		url: g.url + url,
		r: g.r,
	}
}

func (g *group) Handle(method Method, url string, middleware ...HandlerFunc) *router {
	return g.r.Handle(method, g.url + url, append(g.middleware, middleware...)...)
}

func (g *group) Get(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodGet, url, middleware...)
}

func (g *group) Post(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodPost, url, middleware...)
}

func (g *group) Put(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodPut, url, middleware...)
}

func (g *group) Delete(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodDelete, url, middleware...)
}

func (g *group) Patch(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodPatch, url, middleware...)
}

func (g *group) Head(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodHead, url, middleware...)
}

func (g *group) Options(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodOptions, url, middleware...)
}

func (g *group) Connect(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodConnect, url, middleware...)
}

func (g *group) Trace(url string, middleware ...HandlerFunc) *router {
	return g.Handle(MethodTrace, url, middleware...)
}
