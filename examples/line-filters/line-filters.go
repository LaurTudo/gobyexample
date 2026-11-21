// Un _filtru de linii_ (line filter) este un tip
// obișnuit de program care citește input-ul de la stdin,
// îl procesează și apoi afișează un rezultat derivat la
// stdout. `grep` și `sed` sunt filtre de linii
// obișnuite.

// Iată un exemplu de filtru de linii în Go care scrie
// o versiune cu majuscule a întregului text primit. Poți
// folosi acest tipar pentru a-ți scrie propriile filtre
// de linii în Go.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Învelind `os.Stdin`, care nu are buffer, într-un
	// scanner cu buffer, obținem o metodă comodă `Scan`
	// care avansează scanner-ul la următorul token; în
	// scanner-ul implicit, acesta este următoarea linie.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` returnează token-ul curent din input,
		// aici următoarea linie.
		ucl := strings.ToUpper(scanner.Text())

		// Afișăm linia cu majuscule.
		fmt.Println(ucl)
	}

	// Verificăm erorile apărute în timpul `Scan`.
	// Sfârșitul fișierului este așteptat și nu este
	// raportat de `Scan` ca eroare.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
