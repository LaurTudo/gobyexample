// _Slice-urile_ sunt un tip de date important în Go,
// oferind o interfață mai puternică pentru secvențe
// decât array-urile.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Spre deosebire de array-uri, tipul unui slice e dat
	// doar de elementele pe care le conține (nu și de
	// numărul lor). Un slice neinițializat este egal cu
	// nil și are lungimea 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Pentru a crea un slice cu lungime nenulă, folosește
	// funcția predefinită `make`. Aici creăm un slice de
	// `string`-uri cu lungimea `3` (inițial cu valori zero).
	// Implicit, capacitatea unui slice nou este egală cu
	// lungimea lui; dacă știm dinainte că slice-ul va
	// crește, putem transmite explicit o capacitate ca
	// parametru suplimentar pentru `make`.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Putem seta și citi valori la fel ca la array-uri.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returnează lungimea slice-ului, cum era de
	// așteptat.
	fmt.Println("len:", len(s))

	// Pe lângă aceste operații de bază, slice-urile
	// suportă și altele, care le fac mai bogate decât
	// array-urile. Una dintre ele este funcția predefinită
	// `append`, care returnează un slice ce conține una
	// sau mai multe valori noi. Observă că trebuie să
	// preluăm valoarea returnată de `append`, deoarece
	// putem primi o nouă valoare de tip slice.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice-urile pot fi și copiate cu `copy`. Aici creăm
	// un slice gol `c`, cu aceeași lungime ca `s`, și
	// copiem în `c` conținutul lui `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice-urile suportă un operator de „feliere” cu
	// sintaxa `slice[low:high]`. De exemplu, aceasta
	// obține un slice cu elementele `s[2]`, `s[3]` și
	// `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Aceasta ia elementele până la `s[5]` (exclusiv).
	l = s[:5]
	fmt.Println("sl2:", l)

	// Iar aceasta ia elementele începând cu `s[2]`
	// (inclusiv).
	l = s[2:]
	fmt.Println("sl3:", l)

	// Putem și declara și inițializa o variabilă de tip
	// slice pe o singură linie.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Pachetul `slices` conține o serie de funcții
	// utilitare folositoare pentru slice-uri.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slice-urile pot fi compuse în structuri de date
	// multidimensionale. Lungimea slice-urilor interioare
	// poate varia, spre deosebire de array-urile
	// multidimensionale.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
