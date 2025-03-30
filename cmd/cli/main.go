package main

import (
	"fmt"
	mlib "github.com/rf742/zapper/internal/lib"
	"os"
)

// Function to adjust sign of force on p1 from p2
func signadjust(p1, p2 mlib.Point) (fs [2]float64) {
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


func main() {
	fmt.Printf("The charge on an electron is %e\n", mlib.E)
	data := mlib.ReadCsvFile(os.Args[1])
	for i:=0; i<len(data); i++ {
		fmt.Println(data[i])
	}
	p1 := mlib.Point{ 0.3, 0.0, 6.0e-9 }
	p2 := mlib.Point{ 0.0, -0.1, -3.0e-9 }
	p3 := mlib.Point{ 0.0, 0.0, 5.0e-9 }
	points := []mlib.Point{p1,p2,p3} 
	var fx, fy float64
	var adj [2]float64
	var forces [][]float64
	for i, _ := range points {
		forces = append(forces, []float64{0.0,0.0})
		for j, _ := range points {
			if i != j {
				adj = signadjust(points[i],points[j])
				fx, fy = points[i].Coul(points[j])
				forces[i][0] += fx*adj[0]
				forces[i][1] += fy*adj[1]
			}
		}
	}
	fmt.Println(points)
	fmt.Println(forces)
}
