// [_Funcțiile variadice_](https://en.wikipedia.org/wiki/Variadic_function)
// pot fi apelate cu orice număr de argumente finale.
// De exemplu, `fmt.Println` este o funcție variadică
// des folosită.

package main

import "fmt"

// Iată o funcție care primește ca argumente un număr
// arbitrar de valori `int`.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// În interiorul funcției, tipul lui `nums` este
	// echivalent cu `[]int`. Putem apela `len(nums)`,
	// putem itera peste el cu `range` etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Funcțiile variadice pot fi apelate în mod obișnuit,
	// cu argumente individuale.
	sum(1, 2)
	sum(1, 2, 3)

	// Dacă ai deja mai multe argumente într-un slice,
	// le poți transmite unei funcții variadice folosind
	// `func(slice...)`, ca aici.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
