// _range_ iterează peste elementele mai multor
// structuri de date predefinite. Să vedem cum se
// folosește `range` cu câteva dintre structurile de
// date pe care le-am învățat deja.

package main

import "fmt"

func main() {

	// Aici folosim `range` pentru a aduna numerele dintr-un
	// slice. La array-uri funcționează la fel.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` pe array-uri și slice-uri oferă atât
	// indexul, cât și valoarea fiecărui element. Mai sus
	// nu aveam nevoie de index, așa că l-am ignorat cu
	// identificatorul gol `_`. Uneori însă chiar avem
	// nevoie de indecși.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` pe un map iterează peste perechile
	// cheie/valoare.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` poate itera și doar peste cheile unui map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` pe string-uri iterează peste punctele de cod
	// Unicode. Prima valoare este indexul byte-ului de
	// început al `rune`-ului, iar a doua este `rune`-ul
	// însuși. Vezi [String-uri și rune](strings-and-runes)
	// pentru mai multe detalii.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
