// În Go, un _array_ este o secvență numerotată de
// elemente cu o lungime fixă. În codul Go obișnuit,
// [slice-urile](slices) sunt mult mai des folosite;
// array-urile sunt utile în anumite situații speciale.

package main

import "fmt"

func main() {

	// Aici creăm un array `a` care va conține exact
	// 5 valori de tip `int`. Atât tipul elementelor, cât și
	// lungimea fac parte din tipul array-ului. Implicit, un
	// array este inițializat cu valoarea zero, ceea ce
	// pentru `int` înseamnă `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Putem seta o valoare la un anumit index folosind
	// sintaxa `array[index] = value` și putem citi o valoare
	// cu `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Funcția predefinită `len` returnează lungimea
	// unui array.
	fmt.Println("len:", len(a))

	// Folosește această sintaxă pentru a declara și
	// inițializa un array pe o singură linie.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Poți lăsa și compilatorul să numere elementele
	// în locul tău, folosind `...`
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Dacă specifici indexul cu `:`, elementele dintre
	// ele vor fi inițializate cu zero.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Tipurile array sunt unidimensionale, dar poți
	// compune tipuri pentru a construi structuri de date
	// multidimensionale.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Poți și crea și inițializa array-uri
	// multidimensionale dintr-o dată.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
