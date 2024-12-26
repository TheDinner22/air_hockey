package main

// https://github.com/gorilla/websocket/blob/main/examples/echo/server.go

import (
	"log"
	"net/http"

	"github.com/TheDinner22/air_hockey/handlers"
	// "github.com/TheDinner22/air_hockey/game"
	"github.com/google/uuid"
)

func delme(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
    msg := "<span hx-on:htmx:load=\"ws_session_create()\" id=\"uuid\">" + id.String() + "</span>"
	w.Write([]byte(msg))
}

func main() {
	// pending_games := make(map[uuid.UUID]game.GameState)

	http.HandleFunc("/", handlers.Echo)
	http.HandleFunc("/ws", handlers.Ws_handler)
	http.HandleFunc("/session/create", handlers.Session_create) //handlers.Session_create)
	http.HandleFunc("/session/create_uuid", delme)
	http.HandleFunc("/session/join", handlers.Session_join)

	// file server
	fs := http.FileServer(http.Dir("./public/"))
	http.Handle("/public/", http.StripPrefix("/public", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
