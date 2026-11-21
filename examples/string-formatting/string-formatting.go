// Go oferă un suport excelent pentru formatarea
// string-urilor în tradiția `printf`. Iată câteva
// exemple de operații obișnuite de formatare.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go oferă mai multe „verbe” de afișare, concepute
	// pentru a formata valori Go generale. De exemplu,
	// aici afișăm o instanță a struct-ului `point`.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Dacă valoarea este un struct, varianta `%+v` va
	// include numele câmpurilor struct-ului.
	fmt.Printf("struct2: %+v\n", p)

	// Varianta `%#v` afișează o reprezentare a valorii
	// în sintaxă Go, adică fragmentul de cod sursă care
	// ar produce acea valoare.
	fmt.Printf("struct3: %#v\n", p)

	// Pentru a afișa tipul unei valori, folosește `%T`.
	fmt.Printf("type: %T\n", p)

	// Formatarea valorilor booleene este simplă.
	fmt.Printf("bool: %t\n", true)

	// Există multe opțiuni pentru formatarea numerelor
	// întregi. Folosește `%d` pentru formatarea standard,
	// în baza 10.
	fmt.Printf("int: %d\n", 123)

	// Aici afișăm o reprezentare binară.
	fmt.Printf("bin: %b\n", 14)

	// Aici afișăm caracterul corespunzător numărului
	// întreg dat.
	fmt.Printf("char: %c\n", 33)

	// `%x` oferă codificare hexazecimală.
	fmt.Printf("hex: %x\n", 456)

	// Există și mai multe opțiuni de formatare pentru
	// numerele în virgulă mobilă (float). Pentru
	// formatarea zecimală de bază folosește `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` și `%E` formatează float-ul în notație
	// științifică (în variante ușor diferite).
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// Pentru afișarea simplă a string-urilor folosește
	// `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// Pentru a pune string-urile între ghilimele duble,
	// ca în codul sursă Go, folosește `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// La fel ca la numerele întregi văzute mai devreme,
	// `%x` redă string-ul în baza 16, cu două caractere
	// de output pentru fiecare byte de input.
	fmt.Printf("str3: %x\n", "hex this")

	// Pentru a afișa reprezentarea unui pointer,
	// folosește `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// Când formatezi numere, vei vrea adesea să controlezi
	// lățimea și precizia rezultatului. Pentru a
	// specifica lățimea unui număr întreg, pune un număr
	// după `%` în verb. Implicit, rezultatul va fi
	// aliniat la dreapta și completat cu spații.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// Poți specifica și lățimea float-urilor afișate,
	// deși de obicei vei vrea să limitezi în același
	// timp și precizia zecimală, cu sintaxa
	// lățime.precizie (width.precision).
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// Pentru alinierea la stânga, folosește flag-ul `-`.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// S-ar putea să vrei să controlezi lățimea și la
	// formatarea string-urilor, mai ales ca să te asiguri
	// că se aliniază într-un output de tip tabel. Pentru
	// lățime de bază, cu aliniere la dreapta:
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Pentru alinierea la stânga folosește flag-ul `-`,
	// ca la numere.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Până acum am văzut `Printf`, care afișează
	// string-ul formatat în `os.Stdout`. `Sprintf`
	// formatează și returnează un string fără să-l
	// afișeze nicăieri.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Poți formata și afișa în alte `io.Writers` decât
	// `os.Stdout` folosind `Fprintf`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
