// _Funcțiile_ sunt esențiale în Go. Vom învăța despre
// funcții prin câteva exemple diferite.

package main

import "fmt"

// Iată o funcție care primește două valori `int` și
// returnează suma lor ca `int`.
func plus(a int, b int) int {

	// Go cere instrucțiuni `return` explicite, adică nu
	// returnează automat valoarea ultimei expresii.
	return a + b
}

// Când ai mai mulți parametri consecutivi de același
// tip, poți omite numele tipului pentru acești
// parametri, până la ultimul, care declară tipul.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Apelează o funcție exact cum te-ai aștepta, cu
	// `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
