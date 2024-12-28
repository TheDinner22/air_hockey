package vectors

import (
	"math"
)

// Vec2 could represent a point or some vector
type Vec2 struct {
	X int
	Y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) mag() float64 {
	return math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y))
}

func (self Vec2) dot(other Vec2) int {
	return self.X*other.X + self.Y*other.Y
}

func (self Vec2) angle_between(other Vec2) float64 {
	var cos_theta float64 = float64(self.dot(other)) / (self.mag() * other.mag())
	return math.Acos(cos_theta)
}

func (v *Vec2) scale(num int) {
	v.X *= num
	v.Y *= num
}

// theta is in radians where
// (+) direction is counter-clockwise
// (-) direction is clockwise
func (v *Vec2) rotate(theta float64) {
	mag := v.mag()
	new_angle := math.Atan2(float64(v.Y), float64(v.X)) + theta

	// convert from polar to cartesian
	v.Y = int(math.Sin(new_angle) * mag)
	v.X = int(math.Cos(new_angle) * mag)
}

// self will collide with other, self will not lose any energy and other will not move
// other is considered to have infinite mass, really the angle is what determines the collision
func (self *Vec2) collide_with_rigid(other Vec2) {

}
