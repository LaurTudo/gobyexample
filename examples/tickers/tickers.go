// [Timer-ele](timers) sunt pentru situațiile în care vrei
// să faci ceva o singură dată în viitor - _ticker-ele_
// sunt pentru situațiile în care vrei să faci ceva în
// mod repetat, la intervale regulate. Iată un exemplu
// de ticker care bate periodic până când îl oprim.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ticker-ele folosesc un mecanism similar cu cel al
	// timer-elor: un canal pe care sunt trimise valori.
	// Aici vom folosi instrucțiunea predefinită `select`
	// pe canal pentru a aștepta valorile care sosesc la
	// fiecare 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Ticker-ele pot fi oprite la fel ca timer-ele. Odată
	// oprit, un ticker nu mai primește valori pe canalul
	// său. Pe al nostru îl vom opri după 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
