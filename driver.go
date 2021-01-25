package zeta

type Driver interface {
	Option
	Run(addr ...string) error
}
