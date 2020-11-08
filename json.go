package mvc

import "encoding/json"

type JSON interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type defaultJSON struct {}

func DefaultJSON() JSON{
	return defaultJSON{}
}

func (c defaultJSON) Marshal(v interface{}) ([]byte, error){
	return json.Marshal(v)
}

func (c defaultJSON) Unmarshal(data []byte, v interface{}) error{
	return json.Unmarshal(data, v)
}