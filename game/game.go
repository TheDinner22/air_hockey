package game

// TODO the players velocity should be reset somehow
// it's sorta impossible for a player to have v=0 rn...

import (
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/gorilla/websocket"

	"github.com/TheDinner22/air_hockey/vectors"
)

type Sizes struct {
	Canvas_width  int
	Canvas_height int
}

type Player struct {
	Name  string
	Score int
	Pos   vectors.Circle
	vel   vectors.Vec2
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

	// update velocity TODO you can cheat and make this very big by moving mouse off screen
	player.vel.X = float64(x) - player.Pos.Center.X
	player.vel.Y = float64(y) - player.Pos.Center.Y

	// update pos
	player.Pos.Center.X = float64(x)
	player.Pos.Center.Y = float64(y)

}

func NewPlayer(name string, score int, pos vectors.Circle) Player {
	return Player{Name: name, Score: score, Pos: pos}
}

type Puck struct {
	Pos      vectors.Circle
	Velocity vectors.Vec2 // in this case point is more like a vector dx, dy not x, y
}

func NewPuck(pos vectors.Circle, velocity vectors.Vec2) Puck {
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
	gs.Puck.Pos.Center.X = float64(gs.Game_sizes.Canvas_width / 2)
	gs.Puck.Pos.Center.Y = float64(gs.Game_sizes.Canvas_height / 2)
	gs.Puck.Pos.Radius = gs.Game_sizes.Canvas_width / 15

	gs.P1.Pos.Radius = gs.Game_sizes.Canvas_width / 10
	gs.P1.Pos.Center.X = float64(gs.Game_sizes.Canvas_width / 2)
	// TODO add Y

	gs.P2.Pos.Radius = gs.Game_sizes.Canvas_width / 10
	gs.P2.Pos.Center.X = float64(gs.Game_sizes.Canvas_width / 2)
	// TODO add Y
}

// tick is the smallest increment in time, it's one frame
// the puck should move by its velocity and all collsions/scoring should be checked for and handled
func (gs *GameState) tick() {
	// move the puck
	gs.Puck.Pos.Center.Sum(gs.Puck.Velocity)
	gs.Puck.Velocity.Scale(0.97)

	// clamp maybe
	gs.Puck.Pos.Center.X = min(max(gs.Puck.Pos.Center.X, float64(gs.Puck.Pos.Radius)), float64(gs.Game_sizes.Canvas_width-gs.Puck.Pos.Radius))
	gs.Puck.Pos.Center.Y = min(max(gs.Puck.Pos.Center.Y, float64(gs.Puck.Pos.Radius)), float64(gs.Game_sizes.Canvas_height-gs.Puck.Pos.Radius))

	// puck collsions with a wall?
	puck_x := gs.Puck.Pos.Center.X
	puck_y := gs.Puck.Pos.Center.Y
	puck_radius := float64(gs.Puck.Pos.Radius)

	if puck_x-puck_radius <= 0 || puck_x+puck_radius >= float64(gs.Game_sizes.Canvas_width) {
		gs.Puck.Velocity.Collide_with_rigid(vectors.Y_axis())
	}

	if puck_y-puck_radius <= 0 || puck_y+puck_radius >= float64(gs.Game_sizes.Canvas_height) {
		gs.Puck.Velocity.Collide_with_rigid(vectors.X_axis())
	}

	// puck collsions with P1?
	if gs.P1.Pos.Contains(gs.Puck.Pos) {
        // perform collision
        paddle_as_wall := gs.P1.Pos.Center.With_difference(gs.Puck.Pos.Center).Norm()
        gs.Puck.Velocity.Collide_with_moving_rigid(paddle_as_wall, gs.P1.vel)

        // puck CANNOT get stuck in paddle
        gs.Puck.Pos.Move_out_of(gs.P1.Pos)
	}

	// puck collsions with P2?
	if gs.P2.Pos.Contains(gs.Puck.Pos) {
        // perform collision
        paddle_as_wall := gs.P2.Pos.Center.With_difference(gs.Puck.Pos.Center).Norm()
        gs.Puck.Velocity.Collide_with_moving_rigid(paddle_as_wall, gs.P2.vel)

        // puck CANNOT get stuck in paddle
        gs.Puck.Pos.Move_out_of(gs.P2.Pos)
	}
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
	game_state.Puck.Velocity.Y = 2
	game_state.Puck.Velocity.X = 1

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
			game_state.tick()
			game_state.send_state()
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
