// Într-un exemplu [anterior](range-over-built-in-types)
// am văzut cum `for` și `range` permit iterarea peste
// structurile de date de bază. Putem folosi aceeași
// sintaxă și pentru a itera peste valorile primite
// dintr-un canal.

package main

import "fmt"

func main() {

	// Vom itera peste 2 valori din canalul `queue`.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Acest `range` iterează peste fiecare element pe
	// măsură ce este primit din `queue`. Pentru că am
	// închis canalul mai sus cu `close`, iterația se
	// încheie după primirea celor 2 elemente.
	for elem := range queue {
		fmt.Println(elem)
	}
}
