// `//go:embed` este o
// [directivă de compilator](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives)
// care permite programelor să includă fișiere și foldere
// arbitrare în binarul Go, la momentul compilării. Poți
// citi mai multe despre directiva embed
// [aici](https://pkg.go.dev/embed).
package main

// Importăm pachetul `embed`; dacă nu folosești niciun
// identificator exportat din acest pachet, poți face un
// import gol (blank import) cu `_ "embed"`.
import (
	"embed"
)

// Directivele `embed` acceptă căi relative la
// directorul care conține fișierul sursă Go. Această
// directivă include conținutul fișierului în variabila
// `string` care urmează imediat după ea.
//
//go:embed folder/single_file.txt
var fileString string

// Sau includem conținutul fișierului într-un `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Putem include și mai multe fișiere sau chiar foldere,
// folosind wildcard-uri. Pentru asta folosim o variabilă
// de [tipul embed.FS](https://pkg.go.dev/embed#FS), care
// implementează un sistem de fișiere virtual simplu.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Afișăm conținutul lui `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Citim câteva fișiere din folderul inclus.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
