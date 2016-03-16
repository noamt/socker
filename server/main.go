package main

import (
    "flag"
    "go/build"
    "log"
    "net/http"
    "path/filepath"
    "text/template"
)

var chatWindowTemplate *template.Template

func main() {
    log.Println("Starting Socker server...")
    address := flag.String("address", "localhost:8080", "The address of the Socker server")
    flag.Parse()

    chatWindowTemplate = template.Must(template.ParseFiles(filepath.Join(resources(), "index.html")))
    http.HandleFunc("/", chatWindowHandler)

    h := newChatHub()
    go h.run()
    http.Handle("/chat", chatHandler{hub: h})
    err := http.ListenAndServe(*address, nil)
    if err != nil {
        log.Fatal("Failed to start Socker server", err)
    }
}

func resources() string {
    p, err := build.Default.Import("github.com/noamt/socker/server/resources", "", build.FindOnly)
    if err != nil {
        return "."
    }
    return p.Dir
}

func chatWindowHandler(response http.ResponseWriter, request *http.Request) {
    chatWindowTemplate.Execute(response, request.Host)
}
