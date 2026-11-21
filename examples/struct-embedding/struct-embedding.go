// Go suportă _încorporarea_ (embedding) struct-urilor și
// a interfețelor, pentru a exprima o _compunere_ mai
// fluidă a tipurilor. Nu trebuie confundată cu
// [`//go:embed`](embed-directive), o directivă Go
// introdusă în versiunea 1.16+ pentru a încorpora
// fișiere și foldere în binarul aplicației.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Un `container` _încorporează_ un `base`. O
// încorporare arată ca un câmp fără nume.
type container struct {
	base
	str string
}

func main() {

	// Când creăm struct-uri cu literali, trebuie să
	// inițializăm explicit încorporarea; aici tipul
	// încorporat servește drept nume de câmp.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Putem accesa câmpurile lui `base` direct pe `co`,
	// de exemplu `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternativ, putem scrie calea completă folosind
	// numele tipului încorporat.
	fmt.Println("also num:", co.base.num)

	// Deoarece `container` încorporează `base`, metodele
	// lui `base` devin și metode ale lui `container`.
	// Aici apelăm direct pe `co` o metodă încorporată
	// din `base`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Încorporarea struct-urilor cu metode poate fi
	// folosită pentru a transfera implementări de
	// interfețe către alte struct-uri. Aici vedem că un
	// `container` implementează acum interfața `describer`,
	// pentru că încorporează `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
