package main

import (
    "log"
    "net/url"
    "os"
    "flag"
    "github.com/gorilla/websocket"
    "os/exec"
    "strings"
    "github.com/buildkite/terminal"
    "fmt"
)

type Message struct {
    From string `json:"from"`
    To string `json:"to"`
    Content string `json:"content"`
}

func main() {
    server := flag.String("server", "localhost:8080", "The address (HOST:PORT) of the server")
    label := flag.String("id", "local", "This name that identifies this node")

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
	if err != nil {
	    log.Println("WARN: Error while receiving message - ", err)
	    continue
	}
        if m.To == "all" {
	    if m.Content == "report" {
	        c.WriteJSON(&Message{From: *label, To: "all", Content: "reporting"})
	        continue
	    }
	}
	if m.To == *label {
	    out, err := exec.Command("docker", strings.Split(m.Content, " ")...).CombinedOutput()
	    if err != nil {
	        log.Println("Writing error message to response ")
	        c.WriteJSON(&Message{From: *label, To: "client", Content: string(terminal.Render(out))})
	    } else {
		    log.Println(fmt.Sprintf("%s", terminal.Render(out)))
	            c.WriteJSON(&Message{From: *label, To: "client", Content: string(terminal.Render(out))})
	    }
	    continue
    	}
    }

    os.Exit(0)
}
