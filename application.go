package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/util"
)

var addr = flag.String("addr", "192.168.88.238:80", "http service address")

func main() {

	mem.Init()
	flag.Parse()
	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	//Ensures the websocket is usable
	util.AssignSocketConn(c)
	defer c.Close()
	done := make(chan struct{})
	setupHandler(done, c)
	setupSystemInterrupt(done, interrupt, c)
}

func setupHandler(done chan struct{}, c *websocket.Conn) {
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			controller.HandleIncoming(message)
		}
	}()
}

func setupSystemInterrupt(done chan struct{}, interrupt chan os.Signal, c *websocket.Conn) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
