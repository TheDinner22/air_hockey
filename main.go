package main

// https://github.com/gorilla/websocket/blob/main/examples/echo/server.go

import (
	"log"
	"net/http"

    "github.com/TheDinner22/air_hockey/handlers"
)

func main() {
	http.HandleFunc("/", Echo)
	http.HandleFunc("/ws", Ws_handler)
	http.HandleFunc("/session/create", Session_create)
	http.HandleFunc("/session/join", Session_join)

    // file server
    fs := http.FileServer(http.Dir("./public/"))
    http.Handle("/public/", http.StripPrefix("/public", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
