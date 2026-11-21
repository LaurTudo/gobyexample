// [_Limitarea ratei_](https://en.wikipedia.org/wiki/Rate_limiting)
// (rate limiting) este un mecanism important pentru
// controlul utilizării resurselor și menținerea
// calității serviciului. Go suportă elegant limitarea
// ratei cu goroutine-uri, canale și
// [ticker-e](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Mai întâi vom vedea o limitare simplă a ratei. Să
	// presupunem că vrem să limităm procesarea cererilor
	// primite. Vom servi aceste cereri dintr-un canal cu
	// același nume (`requests`).
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Canalul `limiter` va primi o valoare la fiecare
	// 200 de milisecunde. Acesta este regulatorul din
	// schema noastră de limitare a ratei.
	limiter := time.Tick(200 * time.Millisecond)

	// Blocându-ne pe o primire din canalul `limiter`
	// înainte de a servi fiecare cerere, ne limităm la
	// o cerere la fiecare 200 de milisecunde.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Am putea dori să permitem rafale scurte de cereri
	// în schema noastră de limitare, păstrând totuși
	// limita generală a ratei. Putem face asta adăugând
	// un buffer canalului limitator. Canalul
	// `burstyLimiter` va permite rafale de până la 3
	// evenimente.
	burstyLimiter := make(chan time.Time, 3)

	// Umplem canalul pentru a reprezenta rafala permisă.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// La fiecare 200 de milisecunde vom încerca să
	// adăugăm o valoare nouă în `burstyLimiter`, până la
	// limita lui de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Acum simulăm încă 5 cereri primite. Primele 3
	// dintre ele vor beneficia de capacitatea de rafală
	// a lui `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
