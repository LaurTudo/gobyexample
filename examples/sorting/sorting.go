// Pachetul `slices` din Go implementează sortarea pentru
// tipurile predefinite și pentru tipurile definite de
// utilizator. Vom vedea mai întâi sortarea pentru
// tipurile predefinite.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Funcțiile de sortare sunt generice și funcționează
	// pentru orice tip predefinit _ordonat_. Pentru lista
	// tipurilor ordonate, vezi [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// Un exemplu de sortare a unor `int`-uri.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// Putem folosi pachetul `slices` și pentru a verifica
	// dacă un slice este deja sortat.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
