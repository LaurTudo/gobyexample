// Go suportă
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>funcții recursive</em></a>.
// Iată un exemplu clasic.

package main

import "fmt"

// Funcția `fact` se apelează pe ea însăși până când
// ajunge la cazul de bază, `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Și funcțiile anonime pot fi recursive, dar pentru
	// asta trebuie să declari explicit, cu `var`, o
	// variabilă care să rețină funcția înainte ca aceasta
	// să fie definită.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Deoarece `fib` a fost declarată anterior în
		// `main`, Go știe ce funcție să apeleze aici prin
		// `fib`.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
