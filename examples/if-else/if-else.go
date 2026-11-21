// Ramificarea cu `if` și `else` în Go este
// simplă și directă.

package main

import "fmt"

func main() {

	// Iată un exemplu de bază.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Poți avea o instrucțiune `if` fără `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Operatorii logici precum `&&` și `||` sunt adesea
	// utili în condiții.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// O instrucțiune poate preceda condițiile; orice variabile
	// declarate în această instrucțiune sunt disponibile în
	// ramura curentă și în toate ramurile următoare.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Observă că în Go nu ai nevoie de paranteze în jurul
// condițiilor, dar acoladele sunt obligatorii.
