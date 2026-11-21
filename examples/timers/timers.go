// Adesea vrem să executăm cod Go la un moment dat în
// viitor sau în mod repetat, la un anumit interval.
// Facilitățile predefinite _timer_ și _ticker_ din Go
// fac ușoare ambele sarcini. Vom vedea mai întâi
// timer-ele, apoi [ticker-ele](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timer-ele reprezintă un singur eveniment în viitor.
	// Îi spui timer-ului cât vrei să aștepți, iar el
	// oferă un canal care va fi notificat în acel moment.
	// Acest timer va aștepta 2 secunde.
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` se blochează pe canalul `C` al
	// timer-ului până când acesta trimite o valoare care
	// indică faptul că timer-ul s-a declanșat.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Dacă ai fi vrut doar să aștepți, ai fi putut folosi
	// `time.Sleep`. Un motiv pentru care un timer poate
	// fi util este că îl poți anula înainte să se
	// declanșeze. Iată un exemplu.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Îi dăm lui `timer2` suficient timp să se declanșeze,
	// dacă ar fi urmat să o facă, pentru a arăta că este
	// într-adevăr oprit.
	time.Sleep(2 * time.Second)
}
