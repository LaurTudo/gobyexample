// Go oferă un suport amplu pentru momente de timp și
// durate; iată câteva exemple.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Vom începe prin a obține timpul curent.
	now := time.Now()
	p(now)

	// Poți construi un struct `time` specificând anul,
	// luna, ziua etc. Momentele de timp sunt mereu
	// asociate cu o `Location`, adică un fus orar.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Poți extrage diversele componente ale valorii de
	// timp, cum era de așteptat.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Este disponibilă și ziua săptămânii, `Weekday`
	// (de luni până duminică).
	p(then.Weekday())

	// Aceste metode compară două momente de timp,
	// testând dacă primul are loc înainte, după, respectiv
	// în același timp cu al doilea.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Metoda `Sub` returnează o `Duration` care
	// reprezintă intervalul dintre două momente de timp.
	diff := now.Sub(then)
	p(diff)

	// Putem calcula lungimea duratei în diverse unități
	// de măsură.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Poți folosi `Add` pentru a avansa un moment de timp
	// cu o durată dată, sau cu `-` pentru a-l muta înapoi
	// cu o durată.
	p(then.Add(diff))
	p(then.Add(-diff))
}
