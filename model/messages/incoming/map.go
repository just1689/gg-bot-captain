package incoming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

type MessageShareMap struct {
	Map []model.Tile `json:"model" bson:"model"`
	X   int          `json:"x" bson:"x"`
	Z   int          `json:"z" bson:"z"`
}

func BuildMessageMapFromString(b []byte) (item MessageShareMap, err error) {
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&item)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
