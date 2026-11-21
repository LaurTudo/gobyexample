// În timpul execuției unui program, vrem adesea să
// creăm date de care nu mai avem nevoie după ce
// programul se termină. *Fișierele și directoarele
// temporare* sunt utile în acest scop, pentru că nu
// poluează sistemul de fișiere în timp.

package main

import (
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

	// Cel mai simplu mod de a crea un fișier temporar
	// este apelarea lui `os.CreateTemp`. Aceasta creează
	// un fișier *și* îl deschide pentru citire și
	// scriere. Dăm `""` ca prim argument, așa că
	// `os.CreateTemp` va crea fișierul în locația
	// implicită a sistemului nostru de operare.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Afișăm numele fișierului temporar. Pe sistemele de
	// operare bazate pe Unix, directorul va fi probabil
	// `/tmp`. Numele fișierului începe cu prefixul dat ca
	// al doilea argument al lui `os.CreateTemp`, iar
	// restul este ales automat, astfel încât apelurile
	// concurente să creeze mereu nume de fișiere diferite.
	fmt.Println("Temp file name:", f.Name())

	// Ștergem fișierul după ce am terminat. Sistemul de
	// operare va șterge probabil singur fișierele
	// temporare după o vreme, dar e o practică bună să
	// facem asta explicit.
	defer os.Remove(f.Name())

	// Putem scrie niște date în fișier.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Dacă intenționăm să scriem multe fișiere temporare,
	// am putea prefera să creăm un *director* temporar.
	// Argumentele lui `os.MkdirTemp` sunt aceleași ca ale
	// lui `CreateTemp`, dar returnează *numele* unui
	// director, nu un fișier deschis.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Acum putem construi nume de fișiere temporare,
	// punând în fața lor directorul nostru temporar.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
