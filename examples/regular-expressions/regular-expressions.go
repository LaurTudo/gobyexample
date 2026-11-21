// Go oferă suport integrat pentru [expresii regulate](https://en.wikipedia.org/wiki/Regular_expression).
// Iată câteva exemple de operații obișnuite cu regexp-uri
// în Go.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Aici testăm dacă un pattern se potrivește cu un
	// string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Mai sus am folosit direct un pattern de tip string,
	// dar pentru alte operații cu regexp-uri va trebui
	// să compilezi (`Compile`) un struct `Regexp`
	// optimizat.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Aceste struct-uri au multe metode disponibile. Iată
	// un test de potrivire ca cel de mai devreme.
	fmt.Println(r.MatchString("peach"))

	// Aici găsim potrivirea pentru regexp.
	fmt.Println(r.FindString("peach punch"))

	// Și aici găsim prima potrivire, dar returnăm
	// indicii de început și de sfârșit ai potrivirii în
	// locul textului potrivit.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// Variantele `Submatch` includ informații atât despre
	// potrivirile întregului pattern, cât și despre
	// subpotrivirile din interiorul acestora. De exemplu,
	// aici vom primi informații atât pentru `p([a-z]+)ch`,
	// cât și pentru `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similar, aici vom primi informații despre indicii
	// potrivirilor și ai subpotrivirilor.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Variantele `All` ale acestor funcții se aplică
	// tuturor potrivirilor din input, nu doar primei. De
	// exemplu, pentru a găsi toate potrivirile unui
	// regexp:
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Aceste variante `All` sunt disponibile și pentru
	// celelalte funcții văzute mai sus.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Dacă dăm un număr întreg nenegativ ca al doilea
	// argument al acestor funcții, limităm numărul de
	// potriviri.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Exemplele de mai sus aveau argumente de tip string
	// și foloseau nume precum `MatchString`. Putem da și
	// argumente `[]byte`, renunțând la `String` din
	// numele funcției.
	fmt.Println(r.Match([]byte("peach")))

	// Când creezi variabile globale cu expresii regulate,
	// poți folosi varianta `MustCompile` a lui `Compile`.
	// `MustCompile` intră în panică în loc să returneze o
	// eroare, ceea ce o face mai sigură pentru variabile
	// globale.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// Pachetul `regexp` poate fi folosit și pentru a
	// înlocui porțiuni din string-uri cu alte valori.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Varianta `Func` îți permite să transformi textul
	// potrivit cu o funcție dată.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
