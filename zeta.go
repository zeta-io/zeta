package zeta

type Zeta struct {
	r      *router
	d 	   Driver
	components      map[string]interface{}
}

func New(r *router, d Driver, options ...Option) *Zeta{
	z := &Zeta{
		r: r,
		d: d,
		components: map[string]interface{}{},
	}
	return z.Use(options...)
}

func (z *Zeta) Use(options ...Option) *Zeta{
	for _, option := range options{
		option.Option(z)
	}
	return z
}

func (z *Zeta) Driver(driver Driver) *Zeta{
	z.d = driver
	return z
}

func (z *Zeta) Run(addr... string) error{
	if z.d == nil{
		panic("driver not found.")
	}
	for _, mapping := range z.r.mappings{
		z.d.Handle(mapping.method, mapping.url, mapping.middleware...)
	}
	return z.d.Run(addr...)
}