// Uneori vom dori să sortăm o colecție după alt criteriu
// decât ordinea ei naturală. De exemplu, să presupunem
// că vrem să sortăm string-uri după lungime în loc de
// ordinea alfabetică. Iată un exemplu de sortări
// personalizate în Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Implementăm o funcție de comparare pentru
	// lungimile string-urilor. `cmp.Compare` ne ajută
	// aici.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Acum putem apela `slices.SortFunc` cu această
	// funcție de comparare personalizată pentru a sorta
	// `fruits` după lungimea numelui.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// Putem folosi aceeași tehnică pentru a sorta un
	// slice de valori care nu sunt de tipuri predefinite.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// Sortăm `people` după vârstă folosind
	// `slices.SortFunc`.
	//
	// Notă: dacă struct-ul `Person` este mare, poate
	// vrei ca slice-ul să conțină în schimb `*Person` și
	// să adaptezi funcția de sortare în consecință. Dacă
	// ai dubii, fă un [benchmark](testing-and-benchmarking)!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
