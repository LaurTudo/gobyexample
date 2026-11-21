// _Struct-urile_ din Go sunt colecții de câmpuri cu tip.
// Sunt utile pentru a grupa date împreună, formând
// înregistrări.

package main

import "fmt"

// Tipul struct `person` are câmpurile `name` și `age`.
type person struct {
	name string
	age  int
}

// `newPerson` construiește un nou struct `person` cu
// numele dat.
func newPerson(name string) *person {
	// Go este un limbaj cu garbage collector; poți
	// returna în siguranță un pointer la o variabilă
	// locală - aceasta va fi eliberată de garbage
	// collector doar când nu mai există referințe active
	// către ea.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Această sintaxă creează un struct nou.
	fmt.Println(person{"Bob", 20})

	// Poți numi câmpurile când inițializezi un struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Câmpurile omise vor avea valoarea zero.
	fmt.Println(person{name: "Fred"})

	// Prefixul `&` produce un pointer la struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Este idiomatic să încapsulezi crearea unui struct
	// nou în funcții constructor.
	fmt.Println(newPerson("Jon"))

	// Accesează câmpurile unui struct cu punct.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Struct-urile sunt mutabile.
	s.age = 51
	fmt.Println(s)

	// Poți folosi punctul și cu pointeri la struct-uri -
	// pointerii sunt dereferențiați automat.
	sp := &s
	sp.age = 52
	fmt.Println(sp.age)

	// Dacă un tip struct este folosit pentru o singură
	// valoare, nu trebuie să-i dăm un nume. Valoarea poate
	// avea un tip struct anonim. Această tehnică este
	// folosită des în
	// [testele bazate pe tabele](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
