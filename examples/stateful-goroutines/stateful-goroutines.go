// În exemplul anterior am folosit blocare explicită cu
// [mutex-uri](mutexes) pentru a sincroniza accesul la o
// stare partajată între mai multe goroutine-uri. O altă
// opțiune este să folosim mecanismele de sincronizare
// integrate în goroutine-uri și canale pentru a obține
// același rezultat. Această abordare bazată pe canale
// se potrivește cu ideile din Go de a partaja memoria
// prin comunicare și de a avea fiecare bucată de date
// deținută de exact o singură goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// În acest exemplu, starea noastră va fi deținută de o
// singură goroutine. Astfel garantăm că datele nu vor fi
// niciodată corupte prin acces concurent. Pentru a citi
// sau a scrie acea stare, alte goroutine-uri vor trimite
// mesaje goroutine-ei care o deține și vor primi
// răspunsurile corespunzătoare. Aceste `struct`-uri
// `readOp` și `writeOp` încapsulează acele cereri și o
// modalitate prin care goroutine-a care deține starea
// să răspundă.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Ca înainte, vom număra câte operații efectuăm.
	var readOps atomic.Uint64
	var writeOps atomic.Uint64

	// Canalele `reads` și `writes` vor fi folosite de
	// alte goroutine-uri pentru a trimite cereri de
	// citire, respectiv de scriere.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Aceasta este goroutine-a care deține `state`, un
	// map ca în exemplul anterior, dar acum privat
	// goroutine-ei cu stare. Această goroutine face în
	// mod repetat `select` pe canalele `reads` și
	// `writes`, răspunzând cererilor pe măsură ce sosesc.
	// Un răspuns constă în efectuarea operației cerute,
	// urmată de trimiterea unei valori pe canalul de
	// răspuns `resp` pentru a indica succesul (și
	// valoarea dorită, în cazul `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Pornim 100 de goroutine-uri care trimit citiri
	// goroutine-ei care deține starea, prin canalul
	// `reads`. Fiecare citire presupune construirea unui
	// `readOp`, trimiterea lui pe canalul `reads` și apoi
	// primirea rezultatului pe canalul `resp` furnizat.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				readOps.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Pornim și 10 scrieri, folosind o abordare
	// similară.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				writeOps.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Lăsăm goroutine-urile să lucreze o secundă.
	time.Sleep(time.Second)

	// În final, raportăm numărul de operații.
	fmt.Println("readOps:", readOps.Load())
	fmt.Println("writeOps:", writeOps.Load())
}
