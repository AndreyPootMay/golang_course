package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return (math.Pi) * (c.radius) * (c.radius)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g: ", g);
	fmt.Println("g.area: ", g.area());
	fmt.Println("g.perim: ", g.perim());
	fmt.Println("----------------");
}

func main()  {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	fmt.Println("r: ", r);
	fmt.Println("r.area: ", r.area());
	fmt.Println("r.perim: ", r.perim());
	fmt.Println("----------------");

	fmt.Println("c: ", c);
	fmt.Println("c.area: ", c.area());
	fmt.Println("c.perim: ", c.perim());
	fmt.Println("----------------");
}