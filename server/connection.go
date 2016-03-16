package main

import (
    "github.com/gorilla/websocket"
)

type chatConnection struct {
    webSocketConnection *websocket.Conn
    out chan []byte
    hub *chatHub
}

func (c *chatConnection) startReading() {
    for {
        _, message, err := c.webSocketConnection.ReadMessage()
	if err != nil {
	    break
	}
	c.hub.inboundMessages <- message
    }
    c.webSocketConnection.Close()
}

func (c *chatConnection) startWriting() {
    for message := range c.out {
        err := c.webSocketConnection.WriteMessage(websocket.TextMessage, message)
	if err != nil {
	    break
	}
    }
    c.webSocketConnection.Close()
}
