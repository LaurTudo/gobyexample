// O _goroutine_ este un fir de execuție ușor.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Să presupunem că avem un apel de funcție `f(s)`.
	// Iată cum l-am apela în mod obișnuit, rulându-l
	// sincron.
	f("direct")

	// Pentru a apela această funcție într-o goroutine,
	// folosește `go f(s)`. Noua goroutine se va executa
	// concurent cu cea care a apelat-o.
	go f("goroutine")

	// Poți porni o goroutine și pentru apelul unei
	// funcții anonime.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Cele două apeluri de funcții rulează acum asincron,
	// în goroutine-uri separate. Așteptăm să se termine
	// (pentru o abordare mai robustă, folosește un
	// [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
