package model

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Tag int `json:"tag"`
}

func BuildTagFromString(msg string) (Tag, error) {
	decoder := json.NewDecoder(v)
	var item Tag
	err := decoder.Decode(&item)
	if err != nil {
		fmt.Errorf("There was a problem decoding the post message: %s", err.Error())
	}
	return item, err

}
