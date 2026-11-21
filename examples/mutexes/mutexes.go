// În exemplul anterior am văzut cum să gestionăm o stare
// simplă de contor folosind [operații atomice](atomic-counters).
// Pentru stări mai complexe putem folosi un [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// pentru a accesa în siguranță date din mai multe
// goroutine-uri.

package main

import (
	"fmt"
	"sync"
)

// Container conține un map de contoare; deoarece vrem
// să-l actualizăm concurent din mai multe goroutine-uri,
// adăugăm un `Mutex` pentru a sincroniza accesul.
// Observă că mutex-urile nu trebuie copiate, așa că,
// dacă acest `struct` este transmis mai departe, trebuie
// transmis prin pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Blocăm mutex-ul înainte de a accesa `counters`; îl
	// deblocăm la finalul funcției folosind o instrucțiune
	// [defer](defer).
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Observă că valoarea zero a unui mutex poate fi
		// folosită ca atare, deci nu este necesară nicio
		// inițializare aici.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Această funcție incrementează contorul cu numele
	// dat, într-o buclă.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Rulăm concurent mai multe goroutine-uri; observă
	// că toate accesează același `Container`,
	// iar două dintre ele accesează același contor.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Așteptăm ca goroutine-urile să se termine
	wg.Wait()
	fmt.Println(c.counters)
}
