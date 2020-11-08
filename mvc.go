package mvc

type Mvc struct {
	r      *router
	d 	   Driver
	components      map[string]interface{}
}

func New(r *router, d Driver, options ...Option) *Mvc{
	m := &Mvc{
		r: r,
		d: d,
		components: map[string]interface{}{},
	}
	return m.Use(options...)
}

func (m *Mvc) Use(options ...Option) *Mvc{
	for _, option := range options{
		option.Option(m)
	}
	return m
}

func (m *Mvc) Driver(driver Driver) *Mvc{
	m.d = driver
	return m
}

func (m *Mvc) Run(addr... string) error{
	if m.d == nil{
		panic("driver not found.")
	}
	for _, mapping := range m.r.mappings{
		m.d.Handle(mapping.method, mapping.url, mapping.middleware...)
	}
	return m.d.Run(addr...)
}