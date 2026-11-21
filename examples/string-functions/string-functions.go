// Pachetul `strings` din biblioteca standard oferă multe
// funcții utile pentru lucrul cu string-uri. Iată câteva
// exemple care să-ți facă o idee despre pachet.

package main

import (
	"fmt"
	s "strings"
)

// Dăm lui `fmt.Println` un alias cu un nume mai scurt,
// pentru că îl vom folosi des mai jos.
var p = fmt.Println

func main() {

	// Iată o selecție a funcțiilor disponibile în
	// `strings`. Deoarece acestea sunt funcții ale
	// pachetului, nu metode ale obiectului string în
	// sine, trebuie să transmitem string-ul respectiv ca
	// prim argument al funcției. Poți găsi mai multe
	// funcții în documentația pachetului
	// [`strings`](https://pkg.go.dev/strings).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
