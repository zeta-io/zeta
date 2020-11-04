package json

import "encoding/json"

type JSON interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type defaultJson struct {}

func (c defaultJson) Marshal(v interface{}) ([]byte, error){
	return json.Marshal(v)
}

func (c defaultJson) Unmarshal(data []byte, v interface{}) error{
	return json.Unmarshal(data, v)
}

func Default() JSON{
	return defaultJson{}
}