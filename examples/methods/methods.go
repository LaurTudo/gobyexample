// Go suportă _metode_ definite pe tipuri struct.

package main

import "fmt"

type rect struct {
	width, height int
}

// Metoda `area` are _tipul receptorului_ `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Metodele pot fi definite fie pentru receptori de tip
// pointer, fie pentru receptori de tip valoare. Iată un
// exemplu de receptor de tip valoare.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Aici apelăm cele 2 metode definite pentru
	// struct-ul nostru.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go gestionează automat conversia dintre valori și
	// pointeri la apelurile de metode. Poate vrei să
	// folosești un receptor de tip pointer ca să eviți
	// copierea la apelurile de metode sau ca să permiți
	// metodei să modifice struct-ul receptor.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
