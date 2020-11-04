package types

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"reflect"
	"testing"
)

type convertTest struct {
	src interface{}
	t reflect.Type
	want interface{}
	wantErr bool
}

func TestConvert(t *testing.T) {
	tests := []convertTest{
		{src: true, t: reflect.TypeOf(true), want: true},
		{src: "t", t: reflect.TypeOf(true), want: nil, wantErr: true},
		{src: "123.1", t: reflect.TypeOf(123.1), want: 123.1},
		{src: []string{"hello"}, t: reflect.TypeOf([]string{"hello"}), want: []string{"hello"}},
		{src: "[\"123\"]", t: reflect.TypeOf([]string{"123"}), want: []string{"123"}},
		{src: []string{""}, t: reflect.TypeOf([]int{}), want: nil, wantErr: true},
	}
	for _, test := range tests{
		res, err := Convert(test.src, test.t)
		assert.Equal(t, err != nil, test.wantErr)
		assert.Equal(t, res, test.want)
	}
}

func toString(src interface{}) string{
	bs, _ := json.Marshal(src)
	return string(bs)
}