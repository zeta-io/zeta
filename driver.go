package mvc

import (
	"github.com/vectorgo/mvc/http"
)

type Driver interface{
	Option

	Handle(method http.Method, url string, middleware ...HandlerFunc)
}