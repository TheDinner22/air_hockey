package main

// https://github.com/gorilla/websocket/blob/main/examples/echo/server.go

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ws_handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        log.Println(err)
		return
	}

	// now we can use the conn
	defer conn.Close()

	msg_type, msg, err := conn.ReadMessage()
	if err != nil {
        log.Println(err)
		return
	}

	if err := conn.WriteMessage(msg_type, msg); err != nil {
        log.Println(err)
		return
	}

}

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/ws", ws_handler)

    fs := http.FileServer(http.Dir("./public/"))
    http.Handle("/public/", http.StripPrefix("/public", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
