package main

import (
    "github.com/gorilla/websocket"
    "net/http"
)

var socketUpgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

type chatHandler struct {
    hub *chatHub
}

func (cHandler chatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    socket, err := socketUpgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    c := &chatConnection{out: make(chan []byte, 256), webSocketConnection: socket, hub: cHandler.hub}
    c.hub.registerRequests <- c
    defer func() { c.hub.unregisterRequests <- c }()
    go c.startWriting()
    c.startReading()
}
