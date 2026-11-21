// [Variabilele de mediu](https://en.wikipedia.org/wiki/Environment_variable)
// (environment variables) sunt un mecanism universal pentru
// [transmiterea informațiilor de configurare către programele Unix](https://www.12factor.net/config).
// Să vedem cum setăm, citim și listăm variabilele de
// mediu.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Pentru a seta o pereche cheie/valoare, folosește
	// `os.Setenv`. Pentru a obține valoarea unei chei,
	// folosește `os.Getenv`. Aceasta va returna un string
	// gol dacă cheia nu există în mediu.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Folosește `os.Environ` pentru a lista toate
	// perechile cheie/valoare din mediu. Aceasta
	// returnează un slice de string-uri de forma
	// `KEY=value`. Le poți împărți cu `strings.SplitN`
	// pentru a obține cheia și valoarea. Aici afișăm toate
	// cheile.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
