// _Canalele_ (channels) sunt conductele care conectează
// goroutine-uri concurente. Poți trimite valori într-un
// canal dintr-o goroutine și le poți primi în altă
// goroutine.

package main

import "fmt"

func main() {

	// Creează un canal nou cu `make(chan val-type)`.
	// Tipul unui canal este dat de valorile pe care le
	// transportă.
	messages := make(chan string)

	// _Trimite_ o valoare într-un canal folosind sintaxa
	// `channel <-`. Aici trimitem `"ping"` în canalul
	// `messages` creat mai sus, dintr-o goroutine nouă.
	go func() { messages <- "ping" }()

	// Sintaxa `<-channel` _primește_ o valoare din canal.
	// Aici vom primi mesajul `"ping"` trimis mai sus și
	// îl vom afișa.
	msg := <-messages
	fmt.Println(msg)
}
