package main

import (
    "log"
    "net/url"
    "os"
    "flag"
    "github.com/gorilla/websocket"
    "encoding/json"
)

type Message struct {
    From string `json:"from"`
    To string `json:"to"`
    Content string `json:"content"`
}

func main() {
    server := flag.String("server", "localhost:8080", "The address (HOST:PORT) of the server")
//    label := flag.String("id", "local", "This name that identifies this node")

    flag.Parse()

    u := url.URL{Scheme: "ws", Host: *server, Path: "/chat"}
    log.Println("Connecting to " + u.String())

    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("Failed to dial to server" , err)
    }
    defer c.Close()

    for {
        var m Message
        err := c.ReadJSON(&m)
//        _, m, err := c.ReadMessage()
	if err != nil {
	    log.Println("WARN: Error while receiving message - ", err)
	    continue
	}
	val, err := json.Marshal(m)
	log.Println("Received message: " + string(val))
    }

    os.Exit(0)
}
