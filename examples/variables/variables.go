// În Go, _variabilele_ sunt declarate explicit și sunt
// folosite de compilator, de exemplu, pentru a verifica
// dacă tipurile din apelurile de funcții sunt corecte.

package main

import "fmt"

func main() {

	// `var` declară una sau mai multe variabile.
	var a = "initial"
	fmt.Println(a)

	// Poți declara mai multe variabile deodată.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go deduce tipul variabilelor inițializate.
	var d = true
	fmt.Println(d)

	// Variabilele declarate fără o inițializare primesc
	// _valoarea zero_ a tipului lor. De exemplu, valoarea
	// zero pentru un `int` este `0`.
	var e int
	fmt.Println(e)

	// Sintaxa `:=` este o formă prescurtată pentru
	// declararea și inițializarea unei variabile; în acest
	// caz, este echivalentă cu `var f string = "apple"`.
	// Această sintaxă poate fi folosită doar în funcții.
	f := "apple"
	fmt.Println(f)
}
