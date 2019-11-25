package consts

import "errors"

//error const
var(
	HttpRequestTypeNotMatchError = errors.New("http parser need request type of (*http.Request)")
)