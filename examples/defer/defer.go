// _Defer_ este folosit pentru a ne asigura că un apel de
// funcție este efectuat mai târziu în execuția
// programului, de obicei pentru operații de curățare.
// `defer` este folosit adesea acolo unde în alte limbaje
// s-ar folosi, de exemplu, `ensure` și `finally`.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Să presupunem că vrem să creăm un fișier, să scriem
// în el și apoi să-l închidem când am terminat. Iată
// cum am putea face asta cu `defer`.
func main() {

	// Imediat după ce obținem un obiect de tip fișier cu
	// `createFile`, amânăm închiderea acelui fișier cu
	// `closeFile`. Aceasta va fi executată la finalul
	// funcției care o conține (`main`), după ce
	// `writeFile` s-a terminat.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Este important să verificăm erorile la închiderea
	// unui fișier, chiar și într-o funcție amânată.
	if err != nil {
		panic(err)
	}
}
