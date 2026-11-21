// Scrierea fișierelor în Go urmează tipare similare
// celor pe care le-am văzut mai devreme la citire.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Pentru început, iată cum scriem un string (sau pur
	// și simplu bytes) într-un fișier.
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// Pentru scrieri mai granulare, deschide un fișier
	// pentru scriere.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// Este idiomatic să amâni cu defer un `Close`
	// imediat după deschiderea unui fișier.
	defer f.Close()

	// Poți scrie (`Write`) slice-uri de bytes, cum era
	// de așteptat.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Este disponibilă și metoda `WriteString`.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Apelează `Sync` pentru a scrie efectiv datele pe
	// mediul de stocare permanent.
	f.Sync()

	// `bufio` oferă writeri cu buffer, pe lângă readerii
	// cu buffer pe care i-am văzut mai devreme.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Folosește `Flush` pentru a te asigura că toate
	// operațiile din buffer au fost aplicate writer-ului
	// de dedesubt.
	w.Flush()

}
