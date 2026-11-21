// În acest exemplu vom vedea cum să implementăm un
// _worker pool_ (un grup de workeri) folosind
// goroutine-uri și canale.

package main

import (
	"fmt"
	"time"
)

// Iată worker-ul, din care vom rula mai multe instanțe
// concurente. Acești workeri vor primi sarcini pe
// canalul `jobs` și vor trimite rezultatele
// corespunzătoare pe `results`. Vom aștepta câte o
// secundă pentru fiecare sarcină, ca să simulăm o
// operație costisitoare.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Pentru a folosi grupul nostru de workeri, trebuie
	// să le trimitem sarcini și să le colectăm
	// rezultatele. Pentru asta creăm 2 canale.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Aici pornim 3 workeri, inițial blocați pentru că
	// încă nu există sarcini.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Aici trimitem 5 sarcini pe `jobs`, apoi închidem
	// canalul cu `close` pentru a indica că asta e toată
	// munca pe care o avem.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// În final, colectăm toate rezultatele. Asta ne
	// asigură și că goroutine-urile worker s-au terminat.
	// O altă modalitate de a aștepta mai multe
	// goroutine-uri este să folosești un
	// [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
