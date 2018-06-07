package util

import "time"

var KeyBoardChanDown = make(chan string)
var KeyBoardChanUp = make(chan string)

const KeyDownTime = 100

//StartKeyboard ensures the server is aware of key presses
func StartKeyboard() {
	go func() {
		for {
			select {
			case s := <-KeyBoardChanUp:
				SendMessage(BuildKeyUpMessage(s))
				break
			case s := <-KeyBoardChanDown:
				SendMessage(BuildKeyDownMessage(s))
				break
			}
		}
	}()
}

//KeyUp lifts a key
func KeyUp(in string) {
	go func() {
		KeyBoardChanUp <- in
	}()
}

//KeyDown presses a key down
func KeyDown(in string) {
	go func() {
		KeyBoardChanDown <- in
	}()
}

//PressKey press a key for ms milliseconds
func PressKey(key string, ms int64) {
	KeyDown(key)
	time.Sleep(time.Duration(ms))
	KeyUp(key)
}

//PressLeft turns the tank left
func PressLeft() {
	PressKey("A", KeyDownTime)
}

//PressRight turns the tank left
func PressRight() {
	PressKey("D", KeyDownTime)
}

//PressNumber1 presses a power
func PressNumber1() {
	PressKey("1", KeyDownTime)
}

//KeyUpMessage describes a message to the server
type KeyUpMessage struct {
	Conversation string `json:"conversation" bson:"conversation"`
	Key          string `json:"key" bson:"key"`
}

//KeyDownMessage describes a message to the server
type KeyDownMessage struct {
	Conversation string `json:"conversation" bson:"conversation"`
	Key          string `json:"key" bson:"key"`
}

//BuildKeyUpMessage builds a message object
func BuildKeyUpMessage(key string) KeyUpMessage {
	return KeyUpMessage{Conversation: "P_KU", Key: key}
}

//BuildKeyDownMessage builds a message object
func BuildKeyDownMessage(key string) KeyDownMessage {
	return KeyDownMessage{Conversation: "P_KD", Key: key}
}
