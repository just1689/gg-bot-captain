package incoming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

//MessageListOfGames is an object from the server
type MessageListOfGames struct {
	Games []model.Game `json:"games" bson:"games"`
}

//BuildListOfGamesFromString creates the object from bytes
func BuildListOfGamesFromString(msg []byte) (MessageListOfGames, error) {
	r := bytes.NewReader(msg)
	decoder := json.NewDecoder(r)
	var item MessageListOfGames
	err := decoder.Decode(&item)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err
}
