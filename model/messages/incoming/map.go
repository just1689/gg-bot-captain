package incoming

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/util"
	log "github.com/sirupsen/logrus"
)

//MessageShareMap describes a map
type MessageShareMap struct {
	Map []model.Tile `json:"model" bson:"model"`
	X   int          `json:"x" bson:"x"`
	Z   int          `json:"z" bson:"z"`
}

//BuildMessageMapFromString builds a message from bytes
func BuildMessageMapFromString(b []byte) (item MessageShareMap, err error) {
	if err = util.BytesToDecoder(b).Decode(&item); err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
