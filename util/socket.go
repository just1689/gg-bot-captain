package util

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var con *websocket.Conn
var cha chan interface{}

func AssignSocketConn(co *websocket.Conn) {
	con = co
}

func sendItemImp(item interface{}) {
	log.Infoln(fmt.Sprintf("Sending one now"))
	b, err := json.Marshal(item)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error in util.socket: %s", err.Error()))
		return
	}
	err = con.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error in util.socket: %s", err.Error()))
		return
	}

	log.Infoln(fmt.Sprintf("Sending a message over websocket: %s", string(b)))
}

//SendMessage adds a message to the channel
func SendMessage(item interface{}) {
	sendItemImp(item)

}
