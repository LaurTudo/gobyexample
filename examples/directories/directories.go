// Go are mai multe funcții utile pentru lucrul cu
// *directoare* din sistemul de fișiere.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Creăm un subdirector nou în directorul de lucru
	// curent.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Când creăm directoare temporare, e o practică bună
	// să le amânăm ștergerea cu `defer`. `os.RemoveAll`
	// va șterge un întreg arbore de directoare (similar
	// cu `rm -rf`).
	defer os.RemoveAll("subdir")

	// Funcție ajutătoare pentru a crea un fișier nou,
	// gol.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Putem crea o ierarhie de directoare, inclusiv
	// directoarele părinte, cu `MkdirAll`. Este similar
	// cu comanda `mkdir -p` din linia de comandă.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` listează conținutul unui director,
	// returnând un slice de obiecte `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` ne permite să schimbăm directorul de lucru
	// curent, similar cu `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Acum vom vedea conținutul lui `subdir/parent/child`
	// când listăm directorul *curent*.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Ne întoarcem cu `cd` de unde am pornit.
	err = os.Chdir("../../..")
	check(err)

	// Putem parcurge un director și *recursiv*, inclusiv
	// toate subdirectoarele lui. `WalkDir` acceptă o
	// funcție callback care tratează fiecare fișier sau
	// director vizitat.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// `visit` este apelată pentru fiecare fișier sau
// director găsit recursiv de `filepath.WalkDir`.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
