package handlers

import (
    "log"
    "net/http"
    "fmt"
    "html"

	"github.com/gorilla/websocket"
    "github.com/TheDinner22/air_hockey/game"
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

func Session_create(w http.ResponseWriter, r *http.Request) {
    // we expect the request to look a certian way:
    // - name in the query string, thats it
    name := r.URL.Query().Get("name")
    if name == "" {
        w.WriteHeader(400)
        w.Write([]byte("invalid request: Name missing!"))
        return
    }

    // next we try and create a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        log.Fatal(err) // TODO shouldn't be Fatal
		return
	}

    // create game_state with 1 player (it'll be waiting to be setup)
    // TODO actually init this
    game_state := game.GameState{}

    // for now, we just try and make the ws connection
    game_state.P1_conn = conn
    conn = nil // doesn't hurt to be safe
    defer game_state.P1_conn.Close()

    // block until we get some kinda message
	_, msg, err := conn.ReadMessage()
	if err != nil {
        log.Println(err)
		return
	}
    log.Println(string(msg))
}

func Session_join(w http.ResponseWriter, r *http.Request) {}
