package game

import (
	"github.com/gorilla/websocket"
	"math"
)

type Sizes struct {
	Canvas_width  int
	Canvas_height int
}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

type Circle struct {
	center Point
	radius int
}

func NewCircle(center Point, radius int) Circle {
	return Circle{center, radius}
}

func (c Circle) contains(other Circle) bool {
	diff_x := other.center.x - c.center.x
	diff_y := other.center.y - c.center.y
	center_dist := math.Sqrt(float64(diff_x*diff_x) + float64(diff_y*diff_y))

	longest_dist := c.radius + other.radius

	// if the centers are longest_dist aparart (or further), the circles don't overlap
	if float64(longest_dist) > center_dist {
		return true
	} else {
		return false
	}
}

type Player struct {
	name  string
	score int
	pos   Circle
}

func NewPlayer(name string, score int, pos Circle) Player {
	return Player{name, score, pos}
}

type Puck struct {
	pos      Circle
	velocity Point // in this case point is more like a vector dx, dy not x, y
}

func NewPuck(pos Circle, velocity Point) Puck {
	return Puck{pos, velocity}
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

// basic rules for the game
//1. players can move their mouses, this moves their puck
//2. a player's puck must stay on their half of the board
//3. there is a Puck, it has a velocity and must have collision detection and handling
//4. players can score by getting the puck to collide with a certain part of the oponents side of the board
func Start_game(game_state GameState){
    defer game_state.P1_conn.Close()
    defer game_state.P2_conn.Close()

    msg_type, msg, err := game_state.P1_conn.ReadMessage()
    if err != nil {
        return
    }

    err = game_state.P2_conn.WriteMessage(msg_type, msg)
    if err != nil {
        return
    }
}
