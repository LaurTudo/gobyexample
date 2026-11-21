// Go face posibilă _recuperarea_ (recover) după un
// panic, folosind funcția predefinită `recover`. Un
// `recover` poate împiedica un `panic` să oprească
// programul și îl poate lăsa să-și continue execuția.

// Un exemplu în care asta poate fi util: un server
// n-ar vrea să se oprească brusc dacă una dintre
// conexiunile clienților are o eroare critică. În
// schimb, serverul ar vrea să închidă acea conexiune și
// să continue să servească alți clienți. De fapt, exact
// asta face implicit `net/http` din Go pentru serverele
// HTTP.

package main

import "fmt"

// Această funcție intră în panică.
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` trebuie apelat într-o funcție amânată
	// (deferred). Când funcția care o conține intră în
	// panică, defer-ul se va activa, iar un apel `recover`
	// din interiorul lui va prinde panic-ul.
	defer func() {
		if r := recover(); r != nil {
			// Valoarea returnată de `recover` este eroarea
			// generată în apelul către `panic`.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Acest cod nu va rula, pentru că `mayPanic` intră
	// în panică. Execuția lui `main` se oprește în punctul
	// panic-ului și se reia în closure-ul amânat.
	fmt.Println("After mayPanic()")
}
