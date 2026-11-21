// Mecanismul principal de gestionare a stării în Go este
// comunicarea prin canale. Am văzut asta, de exemplu,
// la [worker pools](worker-pools). Există însă și alte
// opțiuni pentru gestionarea stării. Aici vom vedea
// cum folosim pachetul `sync/atomic` pentru _contoare
// atomice_ accesate de mai multe goroutine-uri.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Vom folosi un tip întreg atomic pentru a reprezenta
	// contorul nostru (mereu pozitiv).
	var ops atomic.Uint64

	// Un WaitGroup ne va ajuta să așteptăm ca toate
	// goroutine-urile să-și termine treaba.
	var wg sync.WaitGroup

	// Vom porni 50 de goroutine-uri, fiecare
	// incrementând contorul de exact 1000 de ori.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// Pentru a incrementa atomic contorul folosim `Add`.
				ops.Add(1)
			}
		})
	}

	// Așteptăm până când toate goroutine-urile s-au terminat.
	wg.Wait()

	// Aici nicio goroutine nu mai scrie în 'ops', dar cu
	// `Load` putem citi atomic o valoare în siguranță
	// chiar și în timp ce alte goroutine-uri o
	// actualizează (atomic).
	fmt.Println("ops:", ops.Load())
}
