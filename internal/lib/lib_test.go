package lib

import (
	"testing"
	"math"
)

func TestPoint(t *testing.T) {
	tolerance := 1e-10
	p1 := Point{ 0.3, 0.0, 6.0e-9 }
	p2 := Point{ 0.0, -0.1, -3.0e-9 }
	p3 := Point{ 0.0, 0.0, 5.0e-9 }
	points := []Point{p1,p2,p3} 
	forces := ResolveForce(points)
	got := forces
	want := [][]float64{{1.1986666666666664e-06, -1.6182e-05}, {-1.7980000000000002e-06, -2.9666999999999996e-05}, {2.996666666666667e-06, -1.3484999999999999e-05}}	
	for i, _ := range got {
		for j, y := range got[i] {
			if math.Abs(y - want[i][j]) > tolerance {
				t.Errorf("Error at (%d, %d) -> got %f, want %f",i,j,y, want[i][j])
			}
		}
	}
}
