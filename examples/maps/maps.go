// _Map-urile_ sunt
// [tipul de date asociativ](https://en.wikipedia.org/wiki/Associative_array)
// predefinit din Go (numite uneori _hashes_ sau _dicts_
// în alte limbaje).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// Pentru a crea un map gol, folosește funcția
	// predefinită `make`: `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Setează perechi cheie/valoare folosind sintaxa
	// obișnuită `name[key] = val`.
	m["k1"] = 7
	m["k2"] = 13

	// Afișarea unui map, de exemplu cu `fmt.Println`, va
	// arăta toate perechile lui cheie/valoare.
	fmt.Println("map:", m)

	// Obține valoarea pentru o cheie cu `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Dacă cheia nu există, se returnează
	// [valoarea zero](https://go.dev/ref/spec#The_zero_value)
	// a tipului valorii.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Funcția predefinită `len`, apelată pe un map,
	// returnează numărul de perechi cheie/valoare.
	fmt.Println("len:", len(m))

	// Funcția predefinită `delete` elimină perechi
	// cheie/valoare dintr-un map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Pentru a elimina *toate* perechile cheie/valoare
	// dintr-un map, folosește funcția predefinită `clear`.
	clear(m)
	fmt.Println("map:", m)

	// A doua valoare returnată, opțională, la citirea
	// unei valori dintr-un map indică dacă cheia există
	// în map. Aceasta poate fi folosită pentru a face
	// diferența între cheile lipsă și cheile cu valori
	// zero, precum `0` sau `""`. Aici nu aveam nevoie de
	// valoarea propriu-zisă, așa că am ignorat-o cu
	// _identificatorul gol_ `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Poți și declara și inițializa un map nou pe aceeași
	// linie, cu această sintaxă.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Pachetul `maps` conține o serie de funcții
	// utilitare folositoare pentru map-uri.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
