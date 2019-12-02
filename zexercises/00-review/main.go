package main

import (
	"fmt"
	"math"
)

func main() {
	exercise1()
	exercise2()
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

// dont' forget to declare TYPE if method RETURNS a value !!
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func exercise1() {
	fmt.Println("------exercise 2------")

	s := square{3, 4.5}
	c := circle{3.1}

	info(c)
	info(s)
}

func exercise2() {
	fmt.Println("------exercise 2------")
	// composite literal slice
	s := []int{2, 3, 17, 5}
	fmt.Println(s)
	for i, v := range s {
		fmt.Println(i, v)
	}

	// composite literal map
	m := map[string]int{
		"age":   12,
		"money": 0,
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	p1 := person{"todd", "withers", []string{"bacon", "egg", "toast"}}
	fmt.Println(p1, p1.first)
	for _, v := range p1.favFoods {
		fmt.Println(v)
	}
	p1.walk()

	sed := sedan{vehicle{"brown", 2}, false}
	fmt.Println(sed.doors)

	tru := truck{vehicle{"red", 4}, true}
	report(sed)
	report(tru)
}

type person struct {
	first    string
	last     string
	favFoods []string
}

func (p person) walk() {
	fmt.Println(p.first, `is walking`)
}

type vehicle struct {
	color string
	doors int
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return fmt.Sprintf("I'm a truck. It's %v that I have four wheels", t.fourWheel)
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintf("I'm a sedan. It's %v that I'm luxe", s.luxury)
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}
