// Router
package zeta

type router struct {
	*group

	middleware []HandlerFunc
	mappings   []Mapping
}

type Mapping struct {
	url        string
	method     Method
	middleware []HandlerFunc
}

func (m *Mapping) Url() string{
	return m.url
}

func (m *Mapping) Method() Method{
	return m.method
}

func (m *Mapping) Middleware() []HandlerFunc{
	return m.middleware
}

func (r *router) Option(z *Zeta) {
	z.r = r
}

func (r *router) Middleware() []HandlerFunc{
	return r.middleware
}

func (r *router) Mappings() []Mapping{
	return r.mappings
}

func (r *router) Handle(method Method, url string, middleware ...HandlerFunc) *router {
	r.mappings = append(r.mappings, Mapping{
		method:     method,
		url:        r.url + url,
		middleware: middleware,
	})
	return r
}

func (r *router) Use(middleware ...HandlerFunc) *router {
	r.middleware = append(r.middleware, middleware...)
	return r
}

func (r *router) Group(url string, middleware ...HandlerFunc) *group {
	return &group{
		middleware: middleware,
		url:        url,
		r:          r,
	}
}