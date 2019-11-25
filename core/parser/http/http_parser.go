package http

import (
	"github.com/vectorgo/mvc/core/consts"
	"github.com/vectorgo/mvc/core/model"
	util2 "github.com/vectorgo/mvc/core/parser/util"
	"io/ioutil"
	"net/http"
)

type Parser struct {}

func (*Parser) Parse(request interface{}) (*model.Request, error){
	if httpReq, ok := request.(*http.Request); ok{
		return parseHttpRequest(httpReq)
	}
	return nil, consts.HttpRequestTypeNotMatchError
}

func parseHttpRequest(request *http.Request) (*model.Request, error){
	baseRequest := &model.Request{
		Method: request.Method,
		Url: request.URL.Path,
		Type: model.HTTP,
		HttpRequest: &model.HttpRequest{
			Request: request,
		},
	}

	if len(request.URL.Query()) > 0{
		baseRequest.Parameters = parseHttpRequestQuery(request)
	}

	if util2.HasRequestBody(request.Method){
		requestBody, err := parseHttpRequestBody(request)
		if err != nil{
			return nil, err
		}
		baseRequest.Body = requestBody
	}
	return baseRequest, nil
}

func parseHttpRequestQuery(request *http.Request) map[string][]string{
	parameters := map[string][]string{}
	for key, values := range request.URL.Query(){
		parameters[key] = values
	}
	return parameters
}

func parseHttpRequestBody(request *http.Request) (string, error){
	body, err := ioutil.ReadAll(request.Body)
	if err != nil{
		return "", err
	}
	return string(body), nil
}