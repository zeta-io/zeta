package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/vectorgo/mvc/http"
	"github.com/vectorgo/mvc/util/types"
	"io/ioutil"
	"net/url"
	"reflect"
	"strings"
)

type Values map[string][]string

func (v Values) Get(key string) string{
	if vs, ok := v[key]; ok{
		if len(vs) > 0{
			return vs[0]
		}
	}
	return ""
}

func (v Values) GetArray(key string) []string{
	if vs, ok := v[key]; ok{
		return vs
	}
	return []string{}
}

type requestParamsProcessor struct {
	c *gin.Context
	contentType string

	body string
	forms Values
	queries Values
}

func newRequestParamsProcessor(c *gin.Context) (*requestParamsProcessor, error){
	contentType := contentType(c)
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		panic(err)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	queries := Values{}
	err = parseQuery(queries, c.Request.URL.RawQuery)
	if err != nil{
		return nil, err
	}

	forms := Values{}
	if contentType == string(http.ContentTypePostForm){
		err = parseQuery(forms, string(body))
		if err != nil{
			return nil, err
		}
	}else if contentType == string(http.ContentTypeFormData){
		//TODO parse multipart/form-data
	}

	return &requestParamsProcessor{
		c: c,
		contentType: contentType,
		body: string(body),
		queries: queries,
		forms: forms,
	}, nil
}

func contentType(c *gin.Context) string{
	return strings.TrimSpace(strings.Split(c.ContentType(), ";")[0])
}

func (p *requestParamsProcessor) process(t reflect.Type, source, name string) (interface{}, error){
	switch source {
	case "query":
		return p.processQuery(t, name)
	case "form":
		return p.processFormData(t, name)
	case "body":
		return p.processJson(t, name)
	}
}

func (p *requestParamsProcessor) processQuery(t reflect.Type, name string) (interface{}, error){
	src := interface{}(nil)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice{
		src = p.queries.GetArray(name)
	}else{
		src = p.queries.Get(name)
	}
	return types.Convert(src, t)
}

func (p *requestParamsProcessor) processFormData(t reflect.Type, name string) (interface{}, error){
	src := interface{}(nil)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice{
		src = p.forms.GetArray(name)
	}else{
		src = p.forms.Get(name)
	}
	return types.Convert(src, t)
}

func (p *requestParamsProcessor) processJson(t reflect.Type, name string) (interface{}, error){
	//TODO json path.
	return types.Convert(p.body, t)
}

func parseQuery(values Values, query string) error {
	var err error
	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := queryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		for _, sub := range strings.Split(value, ","){
			value, err1 = queryUnescape(sub)
			if err1 != nil {
				if err == nil {
					err = err1
				}
				continue
			}
			values[key] = append(values[key], value)
		}
	}
	return err
}

func queryUnescape(v string) (string, error){
	return url.QueryUnescape(v)
}