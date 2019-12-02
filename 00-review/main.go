package main

import "fmt"

// new type
type person struct {
	first string
	last  string
}

// composite type
type secretAgent struct {
	person
	ltk bool
}

// polymorphism
type human interface {
	speak()
}

// methods
func (p person) speak() {
	fmt.Println(p.first, p.last, "regular person")
}

func (s secretAgent) speak() {
	fmt.Println(s.first, s.last, "secret agent")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 3, 4, 5}
	fmt.Println(xi)

	p := person{"todd", "withers"}
	m := map[person]int{
		p: 21,
	}
	fmt.Println(m)
	p.speak()

	jb := secretAgent{person{"james", "bond"}, true}
	jb.speak()

	saySomething(jb)
	saySomething(p)
}
