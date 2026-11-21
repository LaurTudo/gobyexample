// _Timeout-urile_ sunt importante pentru programele care
// se conectează la resurse externe sau care, din alte
// motive, trebuie să-și limiteze timpul de execuție.
// Implementarea timeout-urilor în Go este simplă și
// elegantă datorită canalelor și lui `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// În exemplul nostru, să presupunem că executăm un
	// apel extern care își returnează rezultatul pe un
	// canal `c1` după 2s. Observă că acest canal are
	// buffer, deci trimiterea din goroutine nu este
	// blocantă. Acesta este un tipar des folosit pentru a
	// preveni scurgerile de goroutine-uri în cazul în care
	// canalul nu este citit niciodată.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Iată `select`-ul care implementează un timeout.
	// `res := <-c1` așteaptă rezultatul, iar `<-time.After`
	// așteaptă o valoare trimisă după timeout-ul de 1s.
	// Deoarece `select` continuă cu prima primire care
	// este gata, vom intra pe ramura de timeout dacă
	// operația durează mai mult decât 1s, cât e permis.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Dacă permitem un timeout mai lung, de 3s, primirea
	// din `c2` va reuși și vom afișa rezultatul.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
