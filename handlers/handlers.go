package handlers

import (
    "log"
    "net/http"
    "fmt"
    "html"
    "sync"

    "github.com/TheDinner22/air_hockey/game"

	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var waiting_games = make(map[uuid.UUID]game.GameState)
var waiting_games_mutex sync.Mutex // the zero value is an unlocked mutex

func GetUuid(w http.ResponseWriter, r *http.Request) {// TODO html templates are pretty cool...
	id := uuid.New()
    msg := "<span hx-on:htmx:load=\"ws_session_create()\" id=\"uuid\">" + id.String() + "</span>"
	w.Write([]byte(msg))
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
    // - uuid_str in the query string, thats it

    // first we get the uuid string
    uuid_str := r.URL.Query().Get("uuid")
    if uuid_str == "" {
        w.WriteHeader(400)
        w.Write([]byte("invalid request: Name missing!"))
        return
    }

    // then we attempt to parse it
    uuid, err := uuid.Parse(uuid_str)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid request: malformed uuid string!"))
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
    game_state.Game_sizes = game.Sizes{Canvas_width: 200, Canvas_height: 400}
    game_state.P1_conn = conn

    conn = nil // doesn't hurt to be safe

    // if the super rare chance that this uuid already exists happens, crash
    if _, found := waiting_games[uuid]; found {
        panic("uuid already existed?!?!!?")
    }

    // store the incomplete game_state until someone joins and we can start playing
    waiting_games_mutex.Lock()
    waiting_games[uuid] = game_state
    waiting_games_mutex.Unlock()
}

func Session_join(w http.ResponseWriter, r *http.Request) {
    // we expect the request to look a certian way:
    // - uuid_str in the query string, thats it

    // first we get the uuid string
    uuid_str := r.URL.Query().Get("uuid")
    if uuid_str == "" {
        w.WriteHeader(400)
        w.Write([]byte("invalid request: Name missing!"))
        return
    }

    // then we attempt to parse it
    uuid, err := uuid.Parse(uuid_str)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid request: malformed uuid string!"))
        return
    }

    // get/delete the game state
    waiting_games_mutex.Lock()
    game_state, found := waiting_games[uuid]
    delete(waiting_games, uuid)
    waiting_games_mutex.Unlock()

    if !found {
        w.WriteHeader(400)
        w.Write([]byte("invalid request: that session doesn't exist!"))
        return
    }

    // next we try and create a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        log.Fatal(err) // TODO shouldn't be Fatal
		return
	}

    // update game_state with now conn
    game_state.P2_conn = conn

    // start the game!
    go game.Start_game(game_state)
}
