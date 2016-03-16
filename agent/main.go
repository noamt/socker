package main

import (
    "log"
    "net/url"
    "os"
    "flag"
    "github.com/gorilla/websocket"
)

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
    done := make(chan struct{})
    defer close(done)

    for {
        _, message, err := c.ReadMessage()
	if err != nil {
	    log.Fatal("Error while receiving message", err)
	    return
	}
	log.Println("Received message: " + string(message))
    }

    os.Exit(0)
}
