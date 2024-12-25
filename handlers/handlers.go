package handlers

import (
    "log"
    "net/http"
    "fmt"
    "html"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Ws_handler(w http.ResponseWriter, r *http.Request) {
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

func Session_create(w http.ResponseWriter, r *http.Request) {}
func Session_join(w http.ResponseWriter, r *http.Request) {}
