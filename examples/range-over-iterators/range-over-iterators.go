// Începând cu versiunea 1.23, Go a adăugat suport pentru
// [iteratori](https://go.dev/blog/range-functions),
// care ne permit să folosim `range` peste aproape orice!

package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Să ne uităm din nou la tipul `List` din
// [exemplul anterior](generics). Acolo aveam o metodă
// `AllElements` care returna un slice cu toate
// elementele din listă. Cu iteratorii din Go putem face
// asta mai bine - după cum se vede mai jos.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returnează un _iterator_, care în Go este o
// funcție cu o
// [semnătură specială](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Funcția iterator primește ca parametru o altă
		// funcție, numită prin convenție `yield` (dar
		// numele poate fi oricare). Ea va apela `yield`
		// pentru fiecare element peste care vrem să
		// iterăm și va ține cont de valoarea returnată de
		// `yield` pentru o eventuală oprire timpurie.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iterația nu necesită o structură de date în spate și
// nici măcar nu trebuie să fie finită! Iată o funcție
// care returnează un iterator peste numerele Fibonacci:
// rulează cât timp `yield` returnează `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Deoarece `List.All` returnează un iterator, îl putem
	// folosi într-o buclă `range` obișnuită.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Pachete precum [slices](https://pkg.go.dev/slices)
	// au o serie de funcții utile pentru lucrul cu
	// iteratori. De exemplu, `Collect` primește orice
	// iterator și adună toate valorile lui într-un slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Pachetele din biblioteca standard oferă acum și
	// funcții ajutătoare pentru iteratori. De exemplu,
	// `strings.SplitSeq` iterează peste părțile unui
	// string fără să construiască mai întâi un slice cu
	// rezultatele.
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {

		// Când bucla ajunge la `break` sau la un `return`
		// timpuriu, funcția `yield` transmisă iteratorului
		// va returna `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
