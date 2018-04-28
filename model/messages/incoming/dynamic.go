package incoming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

type MessageShareDynamicThings struct {
	Conversation string        `json:"conversation" bson:"conversation"`
	Things       []model.Thing `json:"things" bson:"things"`
}

func BuildMessageShareDynamicThingsFromString(b []byte) (item MessageShareDynamicThings, err error) {

	log.Infoln(fmt.Sprintf(string(b)))
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&item)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
