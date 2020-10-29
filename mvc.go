package mvc

type Mvc struct {
	r  *router
	driver  Driver
}

func Use(options ...Option) *Mvc{
	m := &Mvc{}
	for _, option := range options{
		option.Option(m)
	}
	return m
}

func (m *Mvc) Driver(driver Driver) *Mvc{
	m.driver = driver
	return m
}

func (m *Mvc) Complete() {
	if m.r == nil{
		return
	}
	for _, mapping := range m.r.mappings{
		m.driver.Handle(mapping.method, mapping.url, mapping.middleware...)
	}
}