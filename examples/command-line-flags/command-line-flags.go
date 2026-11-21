// [_Flag-urile din linia de comandă_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// sunt un mod obișnuit de a specifica opțiuni pentru
// programele din linia de comandă. De exemplu, în
// `wc -l`, `-l` este un flag din linia de comandă.

package main

// Go oferă pachetul `flag`, care suportă parsarea de
// bază a flag-urilor din linia de comandă. Vom folosi
// acest pachet pentru a implementa programul nostru de
// exemplu.
import (
	"flag"
	"fmt"
)

func main() {

	// Declarațiile de bază de flag-uri sunt disponibile
	// pentru opțiuni de tip string, întreg și boolean.
	// Aici declarăm un flag string `word` cu valoarea
	// implicită `"foo"` și o scurtă descriere. Funcția
	// `flag.String` returnează un pointer la string (nu o
	// valoare string); vom vedea mai jos cum folosim
	// acest pointer.
	wordPtr := flag.String("word", "foo", "a string")

	// Aici declarăm flag-urile `numb` și `fork`,
	// folosind o abordare similară cu cea pentru `word`.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Se poate declara și o opțiune care folosește o
	// variabilă existentă, declarată în altă parte a
	// programului. Observă că trebuie să transmitem un
	// pointer funcției de declarare a flag-ului.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// După ce toate flag-urile sunt declarate, apelează
	// `flag.Parse()` pentru a parsa linia de comandă.
	flag.Parse()

	// Aici doar afișăm opțiunile parsate și eventualele
	// argumente poziționale de la final. Observă că
	// trebuie să dereferențiem pointerii, de exemplu cu
	// `*wordPtr`, pentru a obține valorile efective ale
	// opțiunilor.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
