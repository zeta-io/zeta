package strings

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

type valueOfTest struct {
	dest interface{}
	want string
}

func TestValueOf(t *testing.T) {
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

	tests := []valueOfTest{
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
		assert.Equal(t, ValueOf(test.dest), test.want)
	}
}