package main

import (
    "github.com/gorilla/websocket"
)

type chatConnection struct {
    webSocketConnection *websocket.Conn
    out chan []byte
    hub *chatHub
}

type message struct {
    from string
    to string
    content string
}

func (c *chatConnection) startReading() {
    for {
        _, m, err := c.webSocketConnection.ReadMessage()
	if err != nil {
	    break
	}
	c.hub.inboundMessages <- m
    }
    c.webSocketConnection.Close()
}

func (c *chatConnection) startWriting() {
    for m := range c.out {
        err := c.webSocketConnection.WriteMessage(websocket.TextMessage, m)
	if err != nil {
	    break
	}
    }
    c.webSocketConnection.Close()
}
