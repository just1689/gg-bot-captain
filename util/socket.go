package util

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
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
			sendAsJson(next, false)
		}
	}()
}

func sendAsJson(item interface{}, verbose bool) {
	b, err := json.Marshal(item)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("Error in util.socket: %s", err.Error()))
		return
	}

	if con == nil {
		logrus.Errorln(fmt.Sprintf("Error in util.socket: Socket is nil"))
		return
	}

	err = con.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("Error in util.socket: %s", err.Error()))
		return
	}

	if verbose {
		logrus.Infoln(fmt.Sprintf("Sending a message over websocket: %s", string(b)))
	}
}

//SendMessage adds a message to the channel
func SendMessage(item interface{}) {
	conChan <- item
}
