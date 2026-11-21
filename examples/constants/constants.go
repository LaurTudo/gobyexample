// Go suportă _constante_ de tip caracter, string, boolean
// și numerice.

package main

import (
	"fmt"
	"math"
)

// `const` declară o valoare constantă.
const s string = "constant"

func main() {
	fmt.Println(s)

	// O declarație `const` poate apărea și în corpul
	// unei funcții.
	const n = 500000000

	// Expresiile constante efectuează calcule aritmetice
	// cu precizie arbitrară.
	const d = 3e20 / n
	fmt.Println(d)

	// O constantă numerică nu are tip până nu i se
	// atribuie unul, de exemplu printr-o conversie
	// explicită.
	fmt.Println(int64(d))

	// Un număr poate primi un tip dacă este folosit într-un
	// context care cere unul, cum ar fi o atribuire către o
	// variabilă sau un apel de funcție. De exemplu, aici
	// `math.Sin` așteaptă un `float64`.
	fmt.Println(math.Sin(n))
}
