/**
2 * @Author: Nico
3 * @Date: 2020/11/8 17:53
4 */
package zeta

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
)

const floatBits = 10000 * 10000 * 10000 * 10000

var (
	uintType        = reflect.TypeOf(uint(0))
	intType         = reflect.TypeOf(0)
	int8Type        = reflect.TypeOf(int8(0))
	uint8Type       = reflect.TypeOf(uint8(0))
	int16Type       = reflect.TypeOf(int16(0))
	uint16Type      = reflect.TypeOf(uint16(0))
	int32Type       = reflect.TypeOf(int32(0))
	uint32Type      = reflect.TypeOf(uint32(0))
	int64Type       = reflect.TypeOf(int64(0))
	uint64Type      = reflect.TypeOf(uint64(0))
	boolType        = reflect.TypeOf(true)
	float32Type     = reflect.TypeOf(float32(0))
	float64Type     = reflect.TypeOf(float64(0))
	stringType      = reflect.TypeOf("")
	stringArrayType = reflect.TypeOf([]string{})
	bytesType       = reflect.TypeOf([]byte{})
)

type Serial interface {
	Serial(dest interface{}) string
	DeSerial(src interface{}, t reflect.Type) (interface{}, error)
}

type defaultSerial struct {
	json JSON
}

func DefaultSerial(jsons ...JSON) Serial {
	var json JSON = defaultJSON{}
	if len(jsons) > 0 {
		json = jsons[0]
	}
	return &defaultSerial{json: json}
}

func (s defaultSerial) Serial(dest interface{}) string {
	var key string
	if dest == nil {
		return key
	}
	switch dest.(type) {
	case float64:
		key = decimal.NewFromFloat(dest.(float64)).String()
	case *float64:
		key = decimal.NewFromFloat(*dest.(*float64)).String()
	case float32:
		key = decimal.NewFromFloat32(dest.(float32)).String()
	case *float32:
		key = decimal.NewFromFloat32(*dest.(*float32)).String()
	case int:
		key = strconv.Itoa(dest.(int))
	case *int:
		key = strconv.Itoa(*dest.(*int))
	case uint:
		key = strconv.Itoa(int(dest.(uint)))
	case *uint:
		key = strconv.Itoa(int(*dest.(*uint)))
	case int8:
		key = strconv.Itoa(int(dest.(int8)))
	case *int8:
		key = strconv.Itoa(int(*dest.(*int8)))
	case uint8:
		key = strconv.Itoa(int(dest.(uint8)))
	case *uint8:
		key = strconv.Itoa(int(*dest.(*uint8)))
	case int16:
		key = strconv.Itoa(int(dest.(int16)))
	case *int16:
		key = strconv.Itoa(int(*dest.(*int16)))
	case uint16:
		key = strconv.Itoa(int(dest.(uint16)))
	case *uint16:
		key = strconv.Itoa(int(*dest.(*uint16)))
	case int32:
		key = strconv.Itoa(int(dest.(int32)))
	case *int32:
		key = strconv.Itoa(int(*dest.(*int32)))
	case uint32:
		key = strconv.Itoa(int(dest.(uint32)))
	case *uint32:
		key = strconv.Itoa(int(*dest.(*uint32)))
	case int64:
		key = strconv.FormatInt(dest.(int64), 10)
	case *int64:
		key = strconv.FormatInt(*dest.(*int64), 10)
	case uint64:
		key = strconv.FormatUint(dest.(uint64), 10)
	case *uint64:
		key = strconv.FormatUint(*dest.(*uint64), 10)
	case string:
		key = dest.(string)
	case *string:
		key = *dest.(*string)
	case []byte:
		key = string(dest.([]byte))
	case *[]byte:
		key = string(*dest.(*[]byte))
	case bool:
		if dest.(bool) {
			key = "true"
		} else {
			key = "false"
		}
	case *bool:
		if *dest.(*bool) {
			key = "true"
		} else {
			key = "false"
		}
	default:
		newValue, _ := s.json.Marshal(dest)
		key = string(newValue)
	}
	return key
}

func (s defaultSerial) DeSerial(src interface{}, t reflect.Type) (interface{}, error) {
	val := s.Serial(src)
	res := interface{}(nil)
	if val == "" && t != stringType {
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
		err := s.json.Unmarshal([]byte(val), &inter)
		if err != nil {
			return nil, err
		}
		res = v.Elem().Interface()
	}
	return res, nil
}
