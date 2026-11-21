// Implicit, canalele sunt _fără buffer_ (unbuffered),
// adică acceptă trimiteri (`chan <-`) doar dacă există
// o primire corespunzătoare (`<- chan`) pregătită să
// preia valoarea trimisă. _Canalele cu buffer_ acceptă
// un număr limitat de valori fără un receptor
// corespunzător pentru ele.

package main

import "fmt"

func main() {

	// Aici creăm cu `make` un canal de string-uri cu un
	// buffer de până la 2 valori.
	messages := make(chan string, 2)

	// Deoarece acest canal are buffer, putem trimite
	// aceste valori în canal fără o primire concurentă
	// corespunzătoare.
	messages <- "buffered"
	messages <- "channel"

	// Mai târziu putem primi aceste două valori în mod
	// obișnuit.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
