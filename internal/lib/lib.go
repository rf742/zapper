package lib

const(
	E = 1.60217663e-19
)

type Point struct {
	X float64
	Y float64
	Q float64
}

/// Method to retrieve position of point
func (p Point) Position() (x, y float64) {
	x = p.X
	y = p.Y
	return
}

/// Method to retrieve charge of point
func (p Point) Charge() (q float64) {
	q = p.Q
	return
}
