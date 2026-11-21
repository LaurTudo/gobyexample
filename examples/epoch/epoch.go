// O cerință frecventă în programe este obținerea
// numărului de secunde, milisecunde sau nanosecunde
// scurse de la [epoca Unix](https://en.wikipedia.org/wiki/Unix_time).
// Iată cum se face asta în Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Folosește `time.Now` cu `Unix`, `UnixMilli` sau
	// `UnixNano` pentru a obține timpul scurs de la epoca
	// Unix în secunde, milisecunde, respectiv
	// nanosecunde.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Poți converti și un număr întreg de secunde sau
	// nanosecunde de la epocă în `time`-ul corespunzător.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
