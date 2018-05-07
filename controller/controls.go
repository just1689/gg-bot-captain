package controller

import (
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/outgoing"
	"github.com/just1689/gg-bot-captain/util"
)

//StartKeyboard ensures the server is aware of key presses
func StartKeyboard() {
	go func() {
		for {
			select {
			case s := <-model.KeyBoardChanUp:
				util.SendMessage(outgoing.BuildKeyUpMessage(s))
				break
			case s := <-model.KeyBoardChanDown:
				util.SendMessage(outgoing.BuildKeyDownMessage(s))
				break
			}
		}
	}()
}

//KeyUp lifts a key
func KeyUp(in string) {
	go func() {
		model.KeyBoardChanUp <- in
	}()
}

//KeyDown presses a key down
func KeyDown(in string) {
	go func() {
		model.KeyBoardChanDown <- in
	}()
}
