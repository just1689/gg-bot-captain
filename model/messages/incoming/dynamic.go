package incoming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

//MessageShareDynamicThings is a wrapper for things
type MessageShareDynamicThings struct {
	Conversation string        `json:"conversation" bson:"conversation"`
	Things       []model.Thing `json:"things" bson:"things"`
}

//BuildMessageShareDynamicThingsFromString creates an object from bytes
func BuildMessageShareDynamicThingsFromString(b []byte) (item MessageShareDynamicThings, err error) {
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&item)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
