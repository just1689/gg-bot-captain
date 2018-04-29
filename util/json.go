package util

import (
	"bytes"
	"encoding/json"
)

func BytesToDecoder(b []byte) *json.Decoder {
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	return decoder
}
