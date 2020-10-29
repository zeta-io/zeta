// Data assembly to struct
package da

import (
	"github.com/vectorgo/mvc/http"
)

const(
	TQuery = "query"
	TBody = "body"
	TRequired = "required"
	TDefault = "default"
)

type DA interface {
	Assembly(url string, body string, contentType http.ContentType, in interface{}) error
}
