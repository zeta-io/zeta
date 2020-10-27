package http

type ContentType string

const(
	ContentTypePostForm ContentType = "application/x-www-form-urlencoded"
	ContentTypeFormData ContentType = "application/form-data"
	ContentTypeMultipart ContentType = "multipart/form-data"
	ContentTypeJSON ContentType = "application/json"
	ContentTypeXML ContentType = "application/xml"
)

