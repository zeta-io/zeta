// Router
package mvc

type router struct {
	url string
	middleware  []HandlerFunc
	mappings    []mapping
}

type mapping struct {
	url        string
	method     Method
	middleware []HandlerFunc
}

func Router(url string, middleware ...HandlerFunc) *router{
	return &router{
		url: url,
		middleware: middleware,
	}
}

func (r *router) Option(m *Mvc){
	m.r = r
}

func (r *router) Use(middleware ...HandlerFunc) *router{
	r.middleware = append(r.middleware, middleware...)
	return r
}

func (r *router) Group(url string, middleware ...HandlerFunc) *group {
	return &group{
		middleware: middleware,
		url: url,
		r: r,
	}
}

func (r *router) Handle(method Method, url string, middleware ...HandlerFunc) *router {
	if r.mappings == nil {
		r.mappings = make([]mapping, 0)
	}
	r.mappings = append(r.mappings, mapping{
		method:     method,
		url:        r.url + url,
		middleware: append(r.middleware, middleware...),
	})
	return r
}

func (r *router) Get(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodGet, url, middleware...)
}

func (r *router) Post(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodPost, url, middleware...)
}

func (r *router) Put(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodPut, url, middleware...)
}

func (r *router) Delete(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodDelete, url, middleware...)
}

func (r *router) Patch(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodPatch, url, middleware...)
}

func (r *router) Head(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodHead, url, middleware...)
}

func (r *router) Options(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodOptions, url, middleware...)
}

func (r *router) Connect(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodConnect, url, middleware...)
}

func (r *router) Trace(url string, middleware ...HandlerFunc) *router {
	return r.Handle(MethodTrace, url, middleware...)
}
