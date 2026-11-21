// Uneori, programele noastre Go trebuie să lanseze alte
// procese.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Vom începe cu o comandă simplă, care nu primește
	// argumente sau input și doar afișează ceva la
	// stdout. Funcția ajutătoare `exec.Command` creează
	// un obiect care reprezintă acest proces extern.
	dateCmd := exec.Command("date")

	// Metoda `Output` rulează comanda, așteaptă să se
	// termine și colectează ieșirea ei standard.
	// Dacă nu au apărut erori, `dateOut` va conține
	// bytes cu informațiile despre dată.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` și alte metode ale lui `Command` vor
	// returna `*exec.Error` dacă a apărut o problemă la
	// executarea comenzii (de exemplu, o cale greșită) și
	// `*exec.ExitError` dacă comanda a rulat, dar s-a
	// terminat cu un cod de ieșire diferit de zero.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			panic(err)
		}
	}

	// În continuare vom vedea un caz puțin mai complicat,
	// în care trimitem prin pipe date către procesul
	// extern, pe `stdin`-ul lui, și colectăm rezultatele
	// de pe `stdout`-ul lui.
	grepCmd := exec.Command("grep", "hello")

	// Aici obținem explicit pipe-urile de input/output,
	// pornim procesul, îi scriem niște input, citim
	// output-ul rezultat și, în final, așteptăm ca
	// procesul să se termine.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Am omis verificările de erori în exemplul de mai
	// sus, dar ai putea folosi tiparul obișnuit
	// `if err != nil` pentru toate. De asemenea, colectăm
	// doar rezultatele din `StdoutPipe`, dar ai putea
	// colecta `StderrPipe` exact în același mod.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Observă că, atunci când lansăm comenzi, trebuie să
	// dăm explicit comanda și array-ul de argumente
	// separat, în loc să putem transmite un singur string
	// de linie de comandă. Dacă vrei să lansezi o comandă
	// completă dintr-un string, poți folosi opțiunea `-c`
	// a lui `bash`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
