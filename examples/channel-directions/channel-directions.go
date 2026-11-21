// Când folosești canale ca parametri de funcție, poți
// specifica dacă un canal este destinat doar trimiterii
// sau doar primirii de valori. Această precizare
// crește siguranța tipurilor în program.

package main

import "fmt"

// Funcția `ping` acceptă doar un canal pentru trimiterea
// de valori. O încercare de a primi valori pe acest
// canal ar produce o eroare la compilare.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Funcția `pong` acceptă un canal pentru primire
// (`pings`) și un al doilea pentru trimitere (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
