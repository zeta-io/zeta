package mvc

type Driver interface{
	Run(addr... string) error
	Handle(method Method, url string, middleware ...HandlerFunc)
}