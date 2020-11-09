/**
2 * @Author: Nico
3 * @Date: 2020/11/8 19:33
4 */
package zeta

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type serialTest struct {
	dest interface{}
	want string
}

type deserialTest struct {
	src interface{}
	t reflect.Type
	want interface{}
	wantErr bool
}

func TestDefaultSerial_DeSerial(t *testing.T) {
	s := defaultSerial{
		json: defaultJSON{},
	}
	tests := []deserialTest{
		{src: true, t: reflect.TypeOf(true), want: true},
		{src: "t", t: reflect.TypeOf(true), want: nil, wantErr: true},
		{src: "123.1", t: reflect.TypeOf(123.1), want: 123.1},
		{src: []string{"hello"}, t: reflect.TypeOf([]string{"hello"}), want: []string{"hello"}},
		{src: "[\"123\"]", t: reflect.TypeOf([]string{"123"}), want: []string{"123"}},
		{src: []string{""}, t: reflect.TypeOf([]int{}), want: nil, wantErr: true},
	}
	for _, test := range tests{
		res, err := s.DeSerial(test.src, test.t)
		assert.Equal(t, err != nil, test.wantErr)
		assert.Equal(t, res, test.want)
	}
}

func TestDefaultSerial_Serial(t *testing.T) {
	s := defaultSerial{
		json: defaultJSON{},
	}
	list := []int{1,2,3}

	int8 := int8(1)
	uint8 := uint8(1)
	int16 := int16(1)
	uint16 := uint16(1)
	int := int(1)
	uint := uint(1)
	int32 := int32(1)
	uint32 := uint32(1)
	int64 := int64(1)
	uint64 := uint64(1)
	float32x1 := float32(1.14)
	float32x2 := float32(1.15)
	float32 := float32(1.1234568)
	float64 := float64(1.123456781234567)
	boolt := true
	boolf := false
	str := "1"
	bs := []byte(str)
	b := byte(1)

	tests := []serialTest{
		{dest: nil, want: ""},
		{dest: b, want: "1"},
		{dest: int8, want: "1"},
		{dest: uint8, want: "1"},
		{dest: int16, want: "1"},
		{dest: uint16, want: "1"},
		{dest: int, want: "1"},
		{dest: uint, want: "1"},
		{dest: int32, want: "1"},
		{dest: uint32, want: "1"},
		{dest: int64, want: "1"},
		{dest: uint64, want: "1"},
		{dest: float32x1, want: "1.14"},
		{dest: float32x2, want: "1.15"},
		{dest: float32, want: "1.1234568"},
		{dest: float64, want: "1.123456781234567"},
		{dest: boolt, want: "true"},
		{dest: boolf, want: "false"},
		{dest: str, want: "1"},
		{dest: bs, want: "1"},
		{dest: list, want: "[1,2,3]"},

		//ptr
		{dest: &b, want: "1"},
		{dest: &int8, want: "1"},
		{dest: &uint8, want: "1"},
		{dest: &int16, want: "1"},
		{dest: &uint16, want: "1"},
		{dest: &int, want: "1"},
		{dest: &uint, want: "1"},
		{dest: &int32, want: "1"},
		{dest: &uint32, want: "1"},
		{dest: &int64, want: "1"},
		{dest: &uint64, want: "1"},
		{dest: &float32x1, want: "1.14"},
		{dest: &float32x2, want: "1.15"},
		{dest: &float32, want: "1.1234568"},
		{dest: &float64, want: "1.123456781234567"},
		{dest: &boolt, want: "true"},
		{dest: &boolf, want: "false"},
		{dest: &str, want: "1"},
		{dest: &bs, want: "1"},
		{dest: &list, want: "[1,2,3]"},
	}

	for _, test := range tests{
		assert.Equal(t, s.Serial(test.dest), test.want)
	}
}
