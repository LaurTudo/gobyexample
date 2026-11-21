// _select_ din Go îți permite să aștepți mai multe
// operații pe canale. Combinarea goroutine-urilor și a
// canalelor cu select este o facilitate puternică a
// limbajului Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// În exemplul nostru vom folosi `select` pe două
	// canale.
	c1 := make(chan string)
	c2 := make(chan string)

	// Fiecare canal va primi o valoare după un anumit
	// timp, pentru a simula, de exemplu, operații RPC
	// blocante executate în goroutine-uri concurente.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Vom folosi `select` pentru a aștepta simultan ambele
	// valori și o vom afișa pe fiecare pe măsură ce
	// sosește.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
