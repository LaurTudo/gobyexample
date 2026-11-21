// În Go este idiomatic ca erorile să fie comunicate
// printr-o valoare returnată explicită, separată. Asta
// contrastează cu excepțiile folosite în limbaje precum
// Java, Python și Ruby și cu valoarea unică, folosită
// atât pentru rezultat, cât și pentru eroare, întâlnită
// uneori în C. Abordarea din Go face ușor de văzut ce
// funcții returnează erori și permite tratarea lor cu
// aceleași construcții ale limbajului folosite și
// pentru celelalte sarcini, care nu țin de erori.
//
// Vezi documentația
// [pachetului errors](https://pkg.go.dev/errors) și
// [acest articol](https://go.dev/blog/go1.13-errors)
// pentru detalii suplimentare.

package main

import (
	"errors"
	"fmt"
)

// Prin convenție, erorile sunt ultima valoare returnată
// și au tipul `error`, o interfață predefinită.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` construiește o valoare `error`
		// simplă, cu mesajul de eroare dat.
		return -1, errors.New("can't work with 42")
	}

	// O valoare `nil` pe poziția erorii indică faptul că
	// nu a existat nicio eroare.
	return arg + 3, nil
}

// O eroare santinelă (sentinel error) este o variabilă
// declarată dinainte, folosită pentru a semnala o
// anumită condiție de eroare.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// Putem împacheta erorile în erori de nivel mai
		// înalt pentru a adăuga context. Cel mai simplu
		// mod de a face asta este cu verbul `%w` din
		// `fmt.Errorf`. Erorile împachetate formează un
		// lanț logic (A împachetează B, care împachetează
		// C etc.), care poate fi interogat cu funcții
		// precum `errors.Is` și `errors.AsType`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// Este idiomatic să verifici eroarea direct pe
		// linia `if`.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` verifică dacă o eroare dată (sau
			// orice eroare din lanțul ei) corespunde unei
			// anumite valori de eroare. Este util mai ales
			// cu erorile împachetate sau imbricate, pentru
			// că îți permite să identifici anumite tipuri
			// de erori sau erori santinelă într-un lanț de
			// erori.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
