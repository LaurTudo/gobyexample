// Pachetul `filepath` oferă funcții pentru a parsa și
// a construi *căi de fișiere* (file paths) într-un mod
// portabil între sistemele de operare; de exemplu,
// `dir/file` pe Linux față de `dir\file` pe Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` ar trebui folosit pentru a construi căi în
	// mod portabil. Primește oricâte argumente și
	// construiește din ele o cale ierarhică.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Ar trebui să folosești mereu `Join` în loc să
	// concatenezi manual `/` sau `\`. Pe lângă
	// portabilitate, `Join` va și normaliza căile,
	// eliminând separatorii de prisos și schimbările
	// de director inutile.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` și `Base` pot fi folosite pentru a împărți o
	// cale în director și fișier. Alternativ, `Split` le
	// returnează pe amândouă în același apel.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Putem verifica dacă o cale este absolută.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Unele nume de fișiere au extensii după un punct.
	// Putem separa extensia din astfel de nume cu `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Pentru a obține numele fișierului fără extensie,
	// folosește `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` găsește o cale relativă între o *bază* (base)
	// și o *țintă* (target). Returnează o eroare dacă
	// ținta nu poate fi exprimată relativ la bază.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
