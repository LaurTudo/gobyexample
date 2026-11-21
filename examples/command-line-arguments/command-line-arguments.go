// [_Argumentele din linia de comandă_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// sunt un mod obișnuit de a parametriza execuția
// programelor. De exemplu, `go run hello.go` transmite
// programului `go` argumentele `run` și `hello.go`.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` oferă acces la argumentele brute din
	// linia de comandă. Observă că prima valoare din
	// acest slice este calea către program, iar
	// `os.Args[1:]` conține argumentele programului.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Poți obține argumente individuale prin indexare
	// obișnuită.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
