// Go suportă [_funcții anonime_](https://en.wikipedia.org/wiki/Anonymous_function),
// care pot forma <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closure-uri</em></a>.
// Funcțiile anonime sunt utile când vrei să definești
// o funcție direct în cod, fără să-i dai un nume.

package main

import "fmt"

// Funcția `intSeq` returnează o altă funcție, pe care
// o definim anonim în corpul lui `intSeq`. Funcția
// returnată _capturează_ variabila `i`, formând astfel
// un closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Apelăm `intSeq` și atribuim rezultatul (o funcție)
	// lui `nextInt`. Această valoare de tip funcție își
	// capturează propria valoare `i`, care va fi
	// actualizată de fiecare dată când apelăm `nextInt`.
	nextInt := intSeq()

	// Observă efectul closure-ului apelând `nextInt`
	// de câteva ori.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Pentru a confirma că starea este proprie acelei
	// funcții, creează și testează una nouă.
	newInts := intSeq()
	fmt.Println(newInts())
}
