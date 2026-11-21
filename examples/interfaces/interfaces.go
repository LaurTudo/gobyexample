// _Interfețele_ sunt colecții denumite de semnături
// de metode.

package main

import (
	"fmt"
	"math"
)

// Iată o interfață de bază pentru forme geometrice.
type geometry interface {
	area() float64
	perim() float64
}

// În exemplul nostru vom implementa această interfață
// pe tipurile `rect` și `circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Pentru a implementa o interfață în Go, trebuie doar
// să implementăm toate metodele din interfață. Aici
// implementăm `geometry` pentru tipul `rect`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implementarea pentru tipul `circle`.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Dacă o variabilă are un tip interfață, putem apela
// metodele din acea interfață. Iată o funcție generică
// `measure` care profită de asta ca să funcționeze cu
// orice `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Uneori e util să știm tipul la runtime al unei valori
// de tip interfață. O opțiune este o *aserțiune de tip*,
// ca aici; alta este un [`switch` pe tipuri](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Tipurile struct `circle` și `rect` implementează
	// amândouă interfața `geometry`, așa că putem folosi
	// instanțe ale acestor struct-uri ca argumente pentru
	// `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
