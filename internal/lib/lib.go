package lib

import "math"

const(
	E = 1.60217663e-19
	K = 8.99e9
)

// Calculates Coulombic force given charges and distance
func CoulombicForce(q1, q2, r float64) float64 {
	if r == 0 {
		return 0
	} else {
		return (K*q1*q2)/(math.Pow(r,2))
	}
}

// Representation of a point charge
type Point struct {
	X float64
	Y float64
	Q float64
}

// Method to retrieve position of point
func (p Point) Position() (x, y float64) {
	x = p.X
	y = p.Y
	return
}

// Method to retrieve charge of point
func (p Point) Charge() (q float64) {
	q = p.Q
	return
}

// Returns the distance between the point and
// another point. 
func (p Point) Dist(other Point) (r float64) {
	xdist := math.Abs(p.X - other.X)
	ydist := math.Abs(p.Y - other.Y)
	r = math.Hypot(xdist, ydist)
	return
}

// Gives back the components of force
// on the point charge from Coulombic
// force given from another charge
func (p Point) Coul(other Point) (fx, fy float64) {
	x := math.Abs(p.X - other.X)
	fx = CoulombicForce(p.Q, other.Q, x)
	y := math.Abs(p.Y - other.Y)
	fy = CoulombicForce(p.Q, other.Q, y)
	return
}

