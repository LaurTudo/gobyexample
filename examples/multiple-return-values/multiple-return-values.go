// Go are suport integrat pentru _valori returnate
// multiple_. Această facilitate este folosită des în
// codul Go idiomatic, de exemplu pentru a returna dintr-o
// funcție atât rezultatul, cât și o eroare.

package main

import "fmt"

// `(int, int)` din semnătura acestei funcții arată că
// funcția returnează 2 valori `int`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Aici folosim cele 2 valori returnate de apel prin
	// _atribuire multiplă_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Dacă vrei doar o parte din valorile returnate,
	// folosește identificatorul gol `_`.
	_, c := vals()
	fmt.Println(c)
}
