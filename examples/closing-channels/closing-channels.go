// _Închiderea_ unui canal indică faptul că nu vor mai fi
// trimise valori pe el. Asta poate fi util pentru a le
// comunica receptorilor canalului că s-a terminat.

package main

import "fmt"

// În acest exemplu vom folosi un canal `jobs` pentru a
// transmite sarcinile de lucru din goroutine-a `main()`
// către o goroutine worker. Când nu mai avem sarcini
// pentru worker, vom închide canalul `jobs` cu `close`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Iată worker-ul, care rulează într-o goroutine.
	// Primește în mod repetat din `jobs` cu
	// `j, more := <-jobs`. În această formă specială de
	// primire, cu 2 valori, `more` va fi `false` dacă
	// `jobs` a fost închis cu `close` și toate valorile
	// din canal au fost deja primite. Folosim asta pentru
	// a trimite o notificare pe `done` când am terminat
	// toate sarcinile.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Aici trimitem 3 sarcini către worker prin canalul
	// `jobs`, apoi îl închidem.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Așteptăm worker-ul folosind metoda de
	// [sincronizare](channel-synchronization) pe care am
	// văzut-o mai devreme.
	<-done

	// Citirea dintr-un canal închis reușește imediat și
	// returnează valoarea zero a tipului de bază. A doua
	// valoare returnată, opțională, este `true` dacă
	// valoarea primită a fost livrată printr-o operație
	// de trimitere reușită în canal, sau `false` dacă a
	// fost o valoare zero generată pentru că canalul este
	// închis și gol.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
