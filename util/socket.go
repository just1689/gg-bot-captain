package util

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var con *websocket.Conn
var conChan = make(chan interface{})

//AssignSocketConn assigns the connection pointer
func AssignSocketConn(co *websocket.Conn) {
	con = co
}

//StartSender starts a goroutine to send messages
func StartSender() {
	go func() {
		for {
			next := <-conChan
			sendItemImp(next, false)
		}
	}()
}

func sendItemImp(item interface{}, verbose bool) {
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

	if verbose {
		log.Infoln(fmt.Sprintf("Sending a message over websocket: %s", string(b)))
	}
}

//SendMessage adds a message to the channel
func SendMessage(item interface{}) {
	conChan <- item
	//sendItemImp(item, false)
}
