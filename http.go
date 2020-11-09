package zeta

type ContentType string
type Method string

const (
	ContentTypePostForm  ContentType = "application/x-www-form-urlencoded"
	ContentTypeFormData ContentType = "multipart/form-data"
	ContentTypeJSON      ContentType = "application/json"
	ContentTypeXML       ContentType = "application/xml"
)

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
	MethodPatch   Method = "PATCH"
)
