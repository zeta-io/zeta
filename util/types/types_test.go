package types

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"reflect"
	"testing"
)

type convertTest struct {
	src interface{}
	want interface{}
}

func TestConvert(t *testing.T) {
	tests := []convertTest{
		{src: true, want: true},
		{src: "123.1", want: 123.1},
		{src: []string{"hello"}, want: []string{"hello"}},
		{src: "[\"123\"]", want: []string{"123"}},
	}
	for _, test := range tests{
		res, err := Convert(test.src, reflect.TypeOf(test.want))
		assert.Equal(t, err, nil)
		assert.Equal(t, res, test.want)
	}
}

func toString(src interface{}) string{
	bs, _ := json.Marshal(src)
	return string(bs)
}