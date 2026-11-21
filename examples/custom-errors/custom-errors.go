// Putem defini tipuri de erori personalizate
// implementând pe ele metoda `Error()`. Iată o variantă
// a exemplului anterior, care folosește un tip
// personalizat pentru a reprezenta explicit o eroare de
// argument.

package main

import (
	"errors"
	"fmt"
)

// Un tip de eroare personalizat are de obicei sufixul
// "Error".
type argError struct {
	arg     int
	message string
}

// Prin adăugarea acestei metode `Error`, `argError`
// implementează interfața `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Returnăm eroarea noastră personalizată.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` este o versiune mai avansată a lui
	// `errors.Is`. Verifică dacă o eroare dată (sau orice
	// eroare din lanțul ei) corespunde unui anumit tip de
	// eroare și o convertește la o valoare de acel tip,
	// returnând și `true`. Dacă nu există potrivire, a
	// doua valoare returnată este `false`.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
