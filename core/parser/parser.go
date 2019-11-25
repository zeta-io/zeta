package parser

import "github.com/vectorgo/mvc/core/model"

//Parser parsing various HTTP requests.
type Parser interface {
	//Parse is the entry point for parsing. Request represents a request,
	//the receiver is the receiver of the request, and the implementer needs
	//to assemble the data in the request to the receiver, if parsing fails,
	//error is returned
	Parse(request interface{}) (*model.Request, error)
}