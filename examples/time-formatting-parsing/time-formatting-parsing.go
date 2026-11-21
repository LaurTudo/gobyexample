// Go suportă formatarea și parsarea timpului prin
// layout-uri bazate pe pattern-uri.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Iată un exemplu simplu de formatare a unui moment
	// de timp conform RFC3339, folosind constanta de
	// layout corespunzătoare.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Parsarea timpului folosește aceleași valori de
	// layout ca `Format`.
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` și `Parse` folosesc layout-uri bazate pe
	// exemple. De obicei vei folosi o constantă din
	// `time` pentru aceste layout-uri, dar poți oferi și
	// layout-uri personalizate. Layout-urile trebuie să
	// folosească momentul de referință
	// `Mon Jan 2 15:04:05 MST 2006` pentru a arăta
	// pattern-ul după care se formatează/parsează un
	// anumit moment de timp/string. Momentul din exemplu
	// trebuie să fie exact cel arătat: anul 2006, 15
	// pentru oră, Monday (luni) pentru ziua săptămânii
	// etc.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// Pentru reprezentări pur numerice poți folosi și
	// formatarea standard a string-urilor, cu
	// componentele extrase din valoarea de timp.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// La un input malformat, `Parse` va returna o eroare
	// care explică problema de parsare.
	_, err := time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM")
	p(err)
}
