// `for` este singura structură de buclă din Go. Iată
// câteva tipuri de bază de bucle `for`.

package main

import "fmt"

func main() {

	// Tipul cel mai simplu, cu o singură condiție.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// O buclă `for` clasică, cu inițializare/condiție/pas.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// O altă modalitate de a realiza iterația simplă „fă
	// asta de N ori” este `range` peste un număr întreg.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` fără condiție se va repeta la nesfârșit până
	// când ieși din buclă cu `break` sau te întorci din
	// funcția care o conține cu `return`.
	for {
		fmt.Println("loop")
		break
	}

	// Poți folosi și `continue` pentru a trece la
	// următoarea iterație a buclei.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
