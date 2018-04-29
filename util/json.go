package util

import (
	"bytes"
	"encoding/json"
)

func bytesToDecoder(b []byte) *json.Decoder {
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	return decoder
}

//Decode codes bytes into a item
func Decode(b []byte, item interface{}) error {
	decoder := bytesToDecoder(b)
	err := decoder.Decode(&item)
	return err
}
