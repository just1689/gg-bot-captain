package incoming

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/util"
	"github.com/sirupsen/logrus"
)

//MessageShareDynamicThings is a wrapper for things
type MessageShareDynamicThings struct {
	Conversation string        `json:"conversation" bson:"conversation"`
	Things       []model.Thing `json:"things" bson:"things"`
}

//BuildMessageShareDynamicThingsFromString creates an object from bytes
func BuildMessageShareDynamicThingsFromString(b []byte) (item MessageShareDynamicThings, err error) {
	if err = util.BytesToDecoder(b).Decode(&item); err != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
