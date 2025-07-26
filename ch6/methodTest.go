package main

import "fmt"
import geometry "example.com/go-demo/ch6/geometry"

func main() {
	twoDis()
	faultTolerant()
	embedding()
	metValAndExp()
}

func twoDis() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 3, Y: 4}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}

func faultTolerant() {
	p := geometry.Point{X: 1, Y: 2}
	(&p).ScaleBy(2)
	p.ScaleBy(2)
	fmt.Println(p.X, p.Y)
}

func embedding() {
	c := geometry.Circle{
    	Point: geometry.Point{X: 1, Y: 2},
    	Radius: 3,
	}
	c.ScaleBy(2)
	fmt.Println(c.X, c.Y, c.Radius)
}

func metValAndExp() {
	p := geometry.Point{X: 0, Y: 0}
	q := geometry.Point{X: 3, Y: 4}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	distanceFunc := geometry.Point.Distance
	fmt.Println(distanceFunc(p, q))
}