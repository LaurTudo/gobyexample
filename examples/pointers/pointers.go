// Go suportă <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointeri</a></em>,
// care îți permit să transmiți referințe către valori și
// înregistrări în cadrul programului.

package main

import "fmt"

// Vom arăta prin contrast cum funcționează pointerii față
// de valori, cu 2 funcții: `zeroval` și `zeroptr`.
// `zeroval` are un parametru `int`, deci argumentele îi
// vor fi transmise prin valoare. `zeroval` va primi o
// copie a lui `ival`, distinctă de cea din funcția
// apelantă.
func zeroval(ival int) {
	ival = 0
}

// În schimb, `zeroptr` are un parametru `*int`, adică
// primește un pointer la `int`. Codul `*iptr` din corpul
// funcției _dereferențiază_ apoi pointerul, trecând de la
// adresa de memorie la valoarea curentă de la acea
// adresă. Atribuirea unei valori unui pointer
// dereferențiat modifică valoarea de la adresa
// referită.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// Sintaxa `&i` oferă adresa de memorie a lui `i`,
	// adică un pointer la `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Și pointerii pot fi afișați.
	fmt.Println("pointer:", &i)

	// Un pointer nou către o valoare poate fi creat cu
	// funcția predefinită `new`.
	p := new(42)
	fmt.Println("value at *p:", *p)
	zeroptr(p)
	fmt.Println("value at *p:", *p)
}
