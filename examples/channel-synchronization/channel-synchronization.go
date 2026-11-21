// Putem folosi canale pentru a sincroniza execuția
// între goroutine-uri. Iată un exemplu în care folosim
// o primire blocantă pentru a aștepta terminarea unei
// goroutine. Când aștepți terminarea mai multor
// goroutine-uri, s-ar putea să preferi un
// [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// Aceasta este funcția pe care o vom rula într-o
// goroutine. Canalul `done` va fi folosit pentru a
// anunța o altă goroutine că funcția și-a terminat
// treaba.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Trimitem o valoare pentru a anunța că am terminat.
	done <- true
}

func main() {

	// Pornim o goroutine worker și îi dăm canalul pe care
	// să trimită notificarea.
	done := make(chan bool, 1)
	go worker(done)

	// Ne blocăm până când primim pe canal o notificare
	// de la worker.
	<-done
}
