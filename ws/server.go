package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8000", nil)
}

type conFds struct {
	sync.Mutex
	data map[time.Duration]websocket.Conn
}

func newConFds() *conFds {
	var cons *conFds

	cons = &conFds{
		data: map[time.Duration]websocket.Conn{},
	}

	go func() {
		for {
			for k, v := range cons.get() {
				err := v.WriteMessage(websocket.PingMessage, []byte(""))
				if err != nil {
					fmt.Println("close")
					v.Close()
					cons.del(k)
				}
			}

			time.Sleep(time.Second * 10)
		}
	}()
	return cons
}

func (c conFds) add(coon websocket.Conn) time.Duration {
	c.Lock()
	defer c.Unlock()

	t := time.Duration(time.Now().UnixNano())
	c.data[t] = coon
	return t
}

func (c conFds) del(t time.Duration) {
	c.Lock()
	defer c.Unlock()

	delete(c.data, t)
}

func (c conFds) get() map[time.Duration]websocket.Conn {
	c.Lock()
	defer c.Unlock()
	var res map[time.Duration]websocket.Conn
	res = c.data
	return res
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	depthCons := newConFds()

	go func() {
		for {
			for k, v := range depthCons.get() {
				err := v.WriteMessage(websocket.TextMessage, []byte("test"))
				if err != nil {
					fmt.Println("close")
					v.Close()
					depthCons.del(k)
				}
			}

			time.Sleep(time.Millisecond * 200)
		}
	}()

	type req struct {
		Sub    string `json:"sub"`
		Symbol string `json:"symbol"`
		Depth  int    `json:"depth"`
	}

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			c.Close()
			break
		}

		var data req
		err = json.Unmarshal(msg, &data)
		if err != nil {
			c.Close()
			break
		}

		err = c.WriteMessage(mt, msg)
		if err != nil {
			c.Close()
			break
		}

		switch data.Sub {
		case "depth":
			fmt.Println(1)
			depthCons.add(*c)
		default:
			c.Close()
			break
		}
	}
}
