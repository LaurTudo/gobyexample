// Un string în Go este un slice de bytes doar pentru
// citire. Limbajul și biblioteca standard tratează
// string-urile în mod special - ca pe niște containere
// de text codificat în [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// În alte limbaje, string-urile sunt formate din
// „caractere”. În Go, conceptul de caracter se numește
// `rune` - este un număr întreg care reprezintă un punct
// de cod Unicode.
// [Acest articol de pe blogul Go](https://go.dev/blog/strings)
// este o introducere bună în subiect.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` este un `string` căruia i se atribuie o valoare
	// literală ce reprezintă cuvântul „salut” în limba
	// thailandeză. Literalii de tip string din Go sunt
	// text codificat în UTF-8.
	const s = "สวัสดี"

	// Deoarece string-urile sunt echivalente cu `[]byte`,
	// aceasta va returna lungimea, în bytes, a datelor
	// brute stocate în string.
	fmt.Println("Len:", len(s))

	// Indexarea unui string produce valorile brute ale
	// byte-ului de la fiecare index. Această buclă
	// generează valorile hexazecimale ale tuturor
	// byte-ilor care alcătuiesc punctele de cod din `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Pentru a număra câte _rune_ conține un string,
	// putem folosi pachetul `utf8`. Observă că timpul de
	// execuție al lui `RuneCountInString` depinde de
	// mărimea string-ului, pentru că trebuie să decodeze
	// secvențial fiecare rună UTF-8. Unele caractere
	// thailandeze sunt reprezentate prin puncte de cod
	// UTF-8 care se pot întinde pe mai mulți bytes, așa
	// că rezultatul acestei numărări poate fi surprinzător.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// O buclă `range` tratează string-urile în mod special
	// și decodează fiecare `rune` împreună cu poziția sa
	// în string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Putem obține aceeași iterație folosind explicit
	// funcția `utf8.DecodeRuneInString`.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Aici demonstrăm transmiterea unei valori `rune`
		// către o funcție.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Valorile încadrate de apostrofuri sunt _literali de
	// tip rune_. Putem compara direct o valoare `rune` cu
	// un literal de tip rune.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
