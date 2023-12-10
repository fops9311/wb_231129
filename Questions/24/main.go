package main

import (
	"fmt"

	p "github.com/fops9311/wb_231129/Questions/24/point"
)

func main() {
	p1 := p.NewPoint(2, -3)
	//_ = p1.x
	p2 := p.NewPoint(-2, -3)
	p3 := p.NewPoint(5, 4)
	fmt.Println(p1.DistanceTo(p2))
	fmt.Println(p2.DistanceTo(p1))
	fmt.Println(p2.DistanceTo(p2))
	fmt.Println(p3.DistanceTo(p2))
	fmt.Println(p3.DistanceTo(p1))
}
