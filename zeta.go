package zeta

import "sync"

type Zeta struct {
	*router
	sync.Mutex

	d          Driver
	loaded       bool
	components map[string]interface{}
}

func New(d Driver, options ...Option) *Zeta {
	router := &router{}
	routerGroup := &group{
		r: router,
	}
	router.group = routerGroup

	z := &Zeta{
		d: d,
		loaded: false,
		router: router,
		components: map[string]interface{}{},
	}
	return z.Options(options...)
}

func (z *Zeta) Options(options ...Option) *Zeta {
	for _, option := range options {
		option.Option(z)
	}
	return z
}

func (z *Zeta) Complete() {
	z.Lock()
	defer z.Unlock()
	if ! z.loaded{
		if z.d == nil {
			panic("driver not found.")
		}
		z.d.Option(z)
		z.loaded = true
	}
}

func (z *Zeta) Run(addr ...string) error {
	z.Complete()
	return z.d.Run(addr...)
}
