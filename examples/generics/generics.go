// Începând cu versiunea 1.18, Go a adăugat suport pentru
// _generice_, cunoscute și ca _parametri de tip_.

package main

import "fmt"

// Ca exemplu de funcție generică, `SlicesIndex` primește
// un slice de orice tip `comparable` și un element de
// acel tip și returnează indexul primei apariții a lui
// v în s, sau -1 dacă nu există. Constrângerea
// `comparable` înseamnă că putem compara valori de acest
// tip cu operatorii `==` și `!=`. Pentru o explicație mai
// detaliată a acestei semnături de tip, vezi
// [acest articol](https://go.dev/blog/deconstructing-type-parameters).
// Observă că această funcție există în biblioteca
// standard ca [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Ca exemplu de tip generic, `List` este o listă simplu
// înlănțuită cu valori de orice tip.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Putem defini metode pe tipuri generice la fel ca pe
// tipurile obișnuite, dar trebuie să păstrăm parametrii
// de tip. Tipul este `List[T]`, nu `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements returnează toate elementele din List ca
// slice. În exemplul următor vom vedea o modalitate mai
// idiomatică de a itera peste toate elementele tipurilor
// personalizate.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// Când apelăm funcții generice, ne putem baza adesea
	// pe _inferența de tip_. Observă că nu trebuie să
	// specificăm tipurile pentru `S` și `E` când apelăm
	// `SlicesIndex` - compilatorul le deduce automat.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... deși le-am putea specifica și explicit.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
