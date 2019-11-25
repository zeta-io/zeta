package model

import "net/http"

type RequestType string

const(
	HTTP RequestType = "HTTP"
	GIN  RequestType= "GIN"
)

type HttpRequest struct {
	Request *http.Request
}

type Request struct {
	Parameters map[string][]string
	Body string
	Method string
	Url string

	Type RequestType
	HttpRequest *HttpRequest `json:"-"`
}