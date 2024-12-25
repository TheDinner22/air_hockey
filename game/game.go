package game

import (
    "math"
)

type Sizes struct {
	canvas_width  int
	canvas_height int
}

type Circle struct {
	x      int
	y      int
	radius int
}

func (c Circle) contains(other Circle) bool { // TODO test this!!!
    diff_x := other.x - c.x
    diff_y := other.y - c.y
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

type Puck struct {
	pos      Circle
	velocity Circle // in this case point is more like a vector dx, dy not x, y
}

type GameState struct {
	p1         Player
	p2         Player
	puck       Puck
	game_sizes Sizes
}
