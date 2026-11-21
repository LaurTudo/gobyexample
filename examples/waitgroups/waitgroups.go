// Pentru a aștepta terminarea mai multor goroutine-uri,
// putem folosi un *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Aceasta este funcția pe care o vom rula în fiecare
// goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Așteptăm, ca să simulăm o operație costisitoare.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Acest WaitGroup este folosit pentru a aștepta
	// terminarea tuturor goroutine-urilor pornite aici.
	// Notă: dacă un WaitGroup este transmis explicit
	// către funcții, trebuie transmis *prin pointer*.
	var wg sync.WaitGroup

	// Pornim mai multe goroutine-uri folosind
	// `WaitGroup.Go`.
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Ne blocăm până când toate goroutine-urile pornite
	// de `wg` s-au terminat. O goroutine s-a terminat
	// când funcția pe care o apelează a returnat.
	wg.Wait()

	// Observă că această abordare nu oferă o modalitate
	// simplă de a propaga erorile de la workeri. Pentru
	// cazuri de utilizare mai avansate, poți folosi
	// [pachetul errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
