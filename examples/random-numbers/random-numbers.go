// Pachetul `math/rand/v2` din Go oferă generarea de
// [numere pseudoaleatoare](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// De exemplu, `rand.IntN` returnează un `int`
	// aleator n, `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` returnează un `float64` `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Asta poate fi folosit pentru a genera float-uri
	// aleatoare în alte intervale, de exemplu
	// `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Dacă vrei un seed cunoscut, creează un nou
	// `rand.Source` și transmite-l constructorului
	// `New`. `NewPCG` creează o nouă sursă
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// care necesită un seed format din două numere
	// `uint64`.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
