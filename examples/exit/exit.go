// Folosește `os.Exit` pentru a ieși imediat din program
// cu un anumit status.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer`-urile _nu_ vor fi rulate când folosim
	// `os.Exit`, așa că acest `fmt.Println` nu va fi
	// apelat niciodată.
	defer fmt.Println("!")

	// Ieșim cu statusul 3.
	os.Exit(3)
}

// Observă că, spre deosebire de C, de exemplu, Go nu
// folosește o valoare întreagă returnată de `main`
// pentru a indica statusul de ieșire. Dacă vrei să ieși
// cu un status diferit de zero, ar trebui să folosești
// `os.Exit`.
