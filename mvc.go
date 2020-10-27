package mvc

import "github.com/vectorgo/mvc/da"

type Mvc struct {
	c       Config
	da      da.DA
	router  *router
}

func New(c Config) *Mvc{
	return &Mvc{
		c: c,
		da: nil,
		router: &router{},
	}
}

func (m *Mvc) DA(da da.DA) *Mvc{
	m.da = da
	return m
}

func (m *Mvc) Router() *router{
	return m.router
}