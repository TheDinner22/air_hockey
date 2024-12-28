package game

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/gorilla/websocket"
)

type Sizes struct {
	Canvas_width  int
	Canvas_height int
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

type Circle struct {
	Center Point
	Radius int
}

func NewCircle(center Point, radius int) Circle {
	return Circle{center, radius}
}

func (c Circle) contains(other Circle) bool {
	diff_x := other.Center.X - c.Center.X
	diff_y := other.Center.Y - c.Center.Y
	center_dist := math.Sqrt(float64(diff_x*diff_x) + float64(diff_y*diff_y))

	longest_dist := c.Radius + other.Radius

	// if the centers are longest_dist aparart (or further), the circles don't overlap
	if float64(longest_dist) > center_dist {
		return true
	} else {
		return false
	}
}

type Player struct {
	Name  string
	Score int
	Pos   Circle
}

// new_pos of the form: "[x, y]"
func (player *Player) update_pos(new_pos *string) {
	no_prefix := strings.TrimPrefix(*new_pos, "[")
	no_postfix := strings.TrimSuffix(no_prefix, "]")
	x_and_y := strings.Split(no_postfix, ",")

	if len(x_and_y) != 2 {
		return
	}

	// TODO err here? idk
	x, err := strconv.Atoi(x_and_y[0])
	if err != nil {
		return
	}
	y, err := strconv.Atoi(x_and_y[1])
	if err != nil {
		return
	}

	player.Pos.Center.X = x
	player.Pos.Center.Y = y

}

func NewPlayer(name string, score int, pos Circle) Player {
	return Player{name, score, pos}
}

type Puck struct {
	Pos      Circle
	Velocity Point // in this case point is more like a vector dx, dy not x, y
}

func NewPuck(pos Circle, velocity Point) Puck {
	return Puck{pos, velocity}
}

func (puck *Puck) tick() {
	puck.Pos.Center.X += puck.Velocity.X
	puck.Pos.Center.Y += puck.Velocity.Y
}

type GameState struct {
	P1         Player
	P2         Player
	Puck       Puck
	Game_sizes Sizes
	P1_conn    *websocket.Conn
	P2_conn    *websocket.Conn
}

func NewGameState(p1 Player, p2 Player, puck Puck, game_sizes Sizes, p1_conn *websocket.Conn, p2_conn *websocket.Conn) GameState {
	return GameState{p1, p2, puck, game_sizes, p1_conn, p2_conn}
}

// I know it takes a ptr, I SWEAR i won't mutate it
func (gs *GameState) send_state() {
	gs_as_json, err := json.Marshal(gs)
	if err != nil {
		panic("gs could not be json!?!?!")
	}

	gs.P1_conn.WriteMessage(websocket.TextMessage, gs_as_json)
	gs.P2_conn.WriteMessage(websocket.TextMessage, gs_as_json)

}

func (gs *GameState) starting_pos() {
	// center the puck
	gs.Puck.Pos.Center.X = gs.Game_sizes.Canvas_width / 2
	gs.Puck.Pos.Center.Y = gs.Game_sizes.Canvas_height / 2
}

// basic rules for the game
// 1. players can move their mouses, this moves their puck
// 2. a player's puck must stay on their half of the board
// 3. there is a Puck, it has a velocity and must have collision detection and handling
// 4. players can score by getting the puck to collide with a certain part of the oponents side of the board
//
// ws stuff:
// 1. either websocket can send us msg's, we should update the gmae_state accordingly
// 2. we send either websocket the game_state (or maybe the updated game state) at the same time
func Start_game(game_state GameState) {
	defer game_state.P1_conn.Close()
	defer game_state.P2_conn.Close()

    // game runs at 60 fps on the server 1/60 is an
    // update every ~17 ms
    // TODO: does this need to be 60 fps???
    ticker := time.NewTicker(time.Millisecond * 17)
    defer ticker.Stop()

	game_state.starting_pos()
    game_state.Puck.Velocity.Y = 1

	// channels for reading
	ch1 := make(chan *string)
	ch2 := make(chan *string)

	go keep_reading(ch1, game_state.P1_conn)
	go keep_reading(ch2, game_state.P2_conn)

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				return
			}
			game_state.P1.update_pos(msg)

		default:
		}

		select {
		case msg, ok := <-ch2:
			if !ok {
				return
			}
			game_state.P2.update_pos(msg)

		default:
		}

		select {
        case <-ticker.C:
            // we only do 60 updates/second ??
            game_state.Puck.tick()
            game_state.send_state()
            fmt.Println("sending!!")
		default:
		}
	}
}

func keep_reading(ch chan *string, conn *websocket.Conn) {
	for {
		msg_type, raw_msg, err := conn.ReadMessage()
		if err != nil {
			close(ch)
			return
		}
		if msg_type == websocket.BinaryMessage {
			close(ch)
			return
		}

		// im told that this doesn't perform a copy...
		msg := string(raw_msg)

		ch <- &msg
	}
}
