package incoming

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/util"
	"github.com/sirupsen/logrus"
)

//MessageListOfGames is an object from the server
type MessageListOfGames struct {
	Games []model.Game `json:"games" bson:"games"`
}

//BuildListOfGamesFromString creates the object from bytes
func BuildListOfGamesFromString(b []byte) (item MessageListOfGames, err error) {
	if err = util.BytesToDecoder(b).Decode(&item); err != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
