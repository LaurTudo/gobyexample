// Citirea și scrierea fișierelor sunt operații de bază
// necesare în multe programe Go. Mai întâi vom vedea
// câteva exemple de citire a fișierelor.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Citirea fișierelor presupune verificarea erorilor la
// majoritatea apelurilor. Această funcție ajutătoare
// ne va simplifica verificările de erori de mai jos.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Poate cea mai simplă operație de citire a unui
	// fișier este încărcarea întregului său conținut în
	// memorie.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// Adesea vei vrea mai mult control asupra modului în
	// care se citește un fișier și a părților citite.
	// Pentru asta, începe prin a deschide (`Open`) un
	// fișier, ca să obții o valoare `os.File`.
	f, err := os.Open(path)
	check(err)

	// Citim câțiva bytes de la începutul fișierului.
	// Permitem citirea a cel mult 5, dar notăm și câți
	// au fost citiți efectiv.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Poți și să te poziționezi (`Seek`) într-un loc
	// cunoscut din fișier și să citești (`Read`) de
	// acolo.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Alte moduri de poziționare sunt relative la
	// poziția curentă a cursorului,
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// și relative la sfârșitul fișierului.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// Pachetul `io` oferă câteva funcții care pot fi
	// utile la citirea fișierelor. De exemplu, citiri ca
	// cele de mai sus pot fi implementate mai robust cu
	// `ReadAtLeast`.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Nu există o funcție predefinită de rewind
	// (revenire la început), dar
	// `Seek(0, io.SeekStart)` face exact asta.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// Pachetul `bufio` implementează un reader cu buffer,
	// util atât pentru eficiența lui în cazul multor
	// citiri mici, cât și pentru metodele de citire
	// suplimentare pe care le oferă.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Închide fișierul când ai terminat (de obicei, asta
	// s-ar programa cu `defer` imediat după deschiderea
	// cu `Open`).
	f.Close()
}
