// Go are diverse tipuri de valori, printre care string-uri,
// numere întregi, numere în virgulă mobilă, valori booleene
// etc. Iată câteva exemple de bază.

package main

import "fmt"

func main() {

	// String-uri, care pot fi concatenate cu `+`.
	fmt.Println("go" + "lang")

	// Numere întregi și numere în virgulă mobilă.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Valori booleene, cu operatorii logici pe care
	// i-ai aștepta.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
