package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8001", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	type req struct {
		Sub    string `json:"sub"`
		Symbol string `json:"symbol"`
		Depth  int    `json:"depth"`
	}

	j := &req{
		Sub:    "depth",
		Symbol: "btc-usdt",
		Depth:  5,
	}

	b, _ := json.Marshal(j)
fmt.Println(string(b))
	c.WriteMessage(websocket.TextMessage, b)

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	select {}
}
