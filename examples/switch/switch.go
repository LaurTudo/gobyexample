// _Instrucțiunile switch_ exprimă condiții cu
// multe ramuri.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Iată un `switch` de bază.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Poți folosi virgule pentru a separa mai multe expresii
	// în aceeași ramură `case`. În acest exemplu folosim și
	// ramura opțională `default`.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` fără expresie este o modalitate alternativă
	// de a exprima logica if/else. Aici arătăm și că
	// expresiile din `case` pot fi neconstante.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Un `switch` pe tipuri compară tipuri în loc de valori.
	// Îl poți folosi pentru a afla tipul unei valori de tip
	// interfață. În acest exemplu, variabila `t` va avea
	// tipul corespunzător ramurii în care se află.
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
