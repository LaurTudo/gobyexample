// _Tipurile enumerate_ (enums) sunt un caz particular de
// [tipuri sumă](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Un enum este un tip care are un număr fix de valori
// posibile, fiecare cu un nume distinct. Go nu are un
// tip enum ca facilitate separată a limbajului, dar
// enum-urile sunt ușor de implementat folosind idiomuri
// existente ale limbajului.

package main

import "fmt"

// Tipul nostru enum `ServerState` are ca tip de bază
// `int`.
type ServerState int

// Valorile posibile pentru `ServerState` sunt definite
// ca constante. Cuvântul-cheie special
// [iota](https://go.dev/ref/spec#Iota) generează automat
// valori constante succesive; în acest caz 0, 1, 2 și
// așa mai departe.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Prin implementarea interfeței
// [fmt.Stringer](https://pkg.go.dev/fmt#Stringer),
// valorile de tip `ServerState` pot fi afișate sau
// convertite în string-uri.
//
// Asta poate deveni anevoios dacă există multe valori
// posibile. În astfel de cazuri, poate fi folosit
// [unealta stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// împreună cu `go:generate` pentru a automatiza
// procesul. Vezi [acest articol](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// pentru o explicație mai detaliată.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Dacă avem o valoare de tip `int`, nu o putem
	// transmite lui `transition` - compilatorul va
	// semnala o nepotrivire de tip. Asta oferă un anumit
	// grad de siguranță a tipurilor la compilare pentru
	// enum-uri.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emulează o tranziție de stare pentru un
// server; primește starea existentă și returnează
// o stare nouă.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Să presupunem că aici verificăm niște condiții
		// pentru a determina starea următoare...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
