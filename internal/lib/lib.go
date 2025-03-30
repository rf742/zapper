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

// Function to adjust sign of force on p1 from p2
func signadjust(p1, p2 Point) (fs [2]float64) {
	fs[0] = 1.0
	fs[1] = 1.0
	// check if both charges are same sign (repulsion)
	if p2.Q*p1.Q > 0{
		// if p1 left rep. force pushes left -> neg
		// if p1 right rep. force pushs right -> pos
		if p2.X - p1.X > 0 {
			fs[0] *= -1
		}
		// if p1 above p2, rep. force pushes up -> pos
		// if p1 below p2, rep. force pushses down -> neg
		
		if p2.Y - p1.Y > 0 {
			fs[1] *= -1
		}
	} else {
		//for attraction, p1 below p2 becomes pos
		if p2.X - p1.X > 0 {
			fs[0] *= -1
		}
		// and p1 left of p2 becomes positive
		if p2.Y - p1.Y > 0 {
			fs[1] *= -1
		}

	}
	return
}

// Gives back the components of force
// on the point charge from Coulombic
// force given from another charge
func (p Point) Coul(other Point) (fx, fy float64) {
	adj := signadjust(p,other) // determine correct sign
	x := math.Abs(p.X - other.X)
	fx = CoulombicForce(p.Q, other.Q, x)
	y := math.Abs(p.Y - other.Y)
	fy = CoulombicForce(p.Q, other.Q, y)
	fx *= adj[0]
	fy *= adj[1]
	return
}

// Accepts a list of points
// Returns a slice of slices. Each interior slice represnts the
// components of the resulting force on that point from all
// other points in the array
func ResolveForce(points []Point) (f [][]float64) {
	var adj [2]float64
	var fx, fy float64
	for i, _ := range points {
		f = append(f, []float64{0.0,0.0})
		for j, _ := range points {
			if i != j {
				adj = signadjust(points[i],points[j])
				fx, fy = points[i].Coul(points[j])
				f[i][0] += fx*adj[0]
				f[i][1] += fy*adj[1]
			}
		}
	}
	return
}	





