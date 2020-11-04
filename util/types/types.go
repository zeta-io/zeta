package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/vectorgo/mvc/util/strings"
	"reflect"
	"strconv"
)

var (
	uintType    = reflect.TypeOf(uint(0))
	intType     = reflect.TypeOf(0)
	int8Type    = reflect.TypeOf(int8(0))
	uint8Type   = reflect.TypeOf(uint8(0))
	int16Type   = reflect.TypeOf(int16(0))
	uint16Type  = reflect.TypeOf(uint64(0))
	int32Type   = reflect.TypeOf(int32(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	int64Type   = reflect.TypeOf(int64(0))
	uint64Type  = reflect.TypeOf(uint64(0))
	boolType    = reflect.TypeOf(true)
	float32Type = reflect.TypeOf(float32(0))
	float64Type = reflect.TypeOf(float64(0))
	stringType  = reflect.TypeOf("")
	stringArrayType = reflect.TypeOf([]string{})
	bytesType   = reflect.TypeOf([]byte{})
)

func Convert(src interface{}, t reflect.Type) (interface{}, error) {
	val := strings.ValueOf(src)
	res := interface{}(nil)
	if val == "" && t != stringType{
		return nil, nil
	}
	switch t {
	case float64Type:
		dec, err := decimal.NewFromString(val)
		if err != nil {
			return nil, err
		}
		res, _ = dec.Float64()
	case float32Type:
		dec, err := decimal.NewFromString(val)
		if err != nil {
			return nil, err
		}
		res, _ = dec.Float64()
	case boolType:
		res = false
		if val == "true" {
			res = true
		} else if val != "false" {
			return nil, errors.New(fmt.Sprintf("%v is not bool type. ", val))
		}
	case intType:
		it, err := strconv.ParseInt(val, 10, strconv.IntSize)
		if err != nil {
			return nil, err
		}
		res = int(it)
	case uintType:
		it, err := strconv.ParseUint(val, 10, strconv.IntSize)
		if err != nil {
			return nil, err
		}
		res = uint(it)
	case int8Type:
		it, err := strconv.ParseInt(val, 10, 8)
		if err != nil {
			return nil, err
		}
		res = int8(it)
	case uint8Type:
		it, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return nil, err
		}
		res = uint8(it)
	case int16Type:
		it, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			return nil, err
		}
		res = int16(it)
	case uint16Type:
		it, err := strconv.ParseUint(val, 10, 16)
		if err != nil {
			return nil, err
		}
		res = uint16(it)
	case int32Type:
		it, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return nil, err
		}
		res = int32(it)
	case uint32Type:
		it, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return nil, err
		}
		res = uint32(it)
	case int64Type:
		it, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		res = it
	case uint64Type:
		it, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return nil, err
		}
		res = it
	case stringType:
		res = val
	case bytesType:
		res = []byte(val)
	default:
		v := reflect.New(t)
		inter := v.Interface()
		err := json.Unmarshal([]byte(val), &inter)
		if err != nil {
			return nil, err
		}
		res = v.Elem().Interface()
	}
	return res, nil
}
