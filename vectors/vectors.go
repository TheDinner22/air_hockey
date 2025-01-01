package vectors

import (
	"math"
)

func X_axis() Vec2 {
	return Vec2{
		X: 1,
		Y: 0,
	}
}

func Y_axis() Vec2 {
	return Vec2{
		X: 0,
		Y: 1,
	}
}

// Vec2 could represent a point or some vector
// Vec2 internally uses floats but only exposes ints so that
// its easy to use (we work on the scale of whole, discrete pixels)
type Vec2 struct {
	// X and Y only exist for json marshalling and should never
	// be written to or read from, their values aren't reliable
	X, Y float64
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) mag() float64 {
	return math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y))
}

func (self Vec2) dot(other Vec2) float64 {
	return self.X*other.X + self.Y*other.Y
}

func (v Vec2) Norm() Vec2 {
	return Vec2{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vec2) Unit_norm() Vec2 {
	norm := v.Norm()
	norm.Scale(1/v.mag())
	return norm
}

func (v *Vec2) Scale(num float64) {
	v.X *= num
	v.Y *= num
}

func (v Vec2) with_scale(num float64) Vec2 {
	return Vec2{
		X: v.X * num,
		Y: v.Y * num,
	}
}

func (v Vec2) With_difference(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (self *Vec2) Sum(other Vec2){
    self.X += other.X
    self.Y += other.Y
}

func (self Vec2) With_sum(other Vec2) Vec2{
    return Vec2{
        X: self.X + other.X,
        Y: self.Y + other.Y,
    }
}

// self will collide with other, self will not lose any energy and other will not move
// other is considered to have infinite mass, really the angle is what determines the collision
func (self *Vec2) Collide_with_rigid(other Vec2) {
	Unit_norm := other.Unit_norm()

	lhs := Unit_norm.with_scale(-1 * self.dot(Unit_norm))
	rhs := self.With_difference(Unit_norm.with_scale(self.dot(Unit_norm)))

    *self = lhs.With_sum(rhs)
}

// note that this method assumes that every shape involved is a circle
func (self *Vec2) Collide_with_moving_rigid(wall Vec2, wall_vel Vec2) {
    // relative velocity is magic!!
    self.X -= wall_vel.X
    self.Y -= wall_vel.Y

    // perform collision as if it were with a stationary wall
    self.Collide_with_rigid(wall)

    self.X += wall_vel.X
    self.Y += wall_vel.Y
}

type Circle struct {
	Center Vec2
	Radius int
}

func NewCircle(center Vec2, radius int) Circle {
	return Circle{center, radius}
}

func (c Circle) Contains(other Circle) bool {
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

func(self *Circle) Move_out_of(other Circle){
    if !other.Contains(*self){
        return
    }

    min_acceptable_dist := self.Radius + other.Radius
    actual_dist := self.Center.With_difference(other.Center).mag()

    direction_to_move := self.Center.With_difference(other.Center).with_scale(1/actual_dist)
    direction_to_move.Scale(float64(min_acceptable_dist)-actual_dist)

    self.Center.Sum(direction_to_move)
}
