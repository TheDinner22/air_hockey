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

// NOT UNIT LENGTH
func (v Vec2) norm() Vec2 {
	return Vec2{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vec2) Unit_norm() Vec2 {
	norm := v.norm()
	norm.scale(v.mag())
	return norm
}

func (v *Vec2) scale(num float64) {
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

func sum(a Vec2, b Vec2) Vec2{
    return Vec2{
        X: a.X + b.X,
        Y: a.Y + b.Y,
    }
}

// self will collide with other, self will not lose any energy and other will not move
// other is considered to have infinite mass, really the angle is what determines the collision
func (self *Vec2) Collide_with_rigid(other Vec2) {
	unit_norm := other.Unit_norm()

	lhs := unit_norm.with_scale(-1 * self.dot(unit_norm))
	rhs := self.With_difference(unit_norm.with_scale(self.dot(unit_norm)))

    *self = sum(lhs, rhs)
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
