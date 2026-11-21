// Biblioteca standard Go oferă unelte simple pentru
// scrierea de loguri din programele Go: pachetul
// [log](https://pkg.go.dev/log) pentru output liber și
// pachetul [log/slog](https://pkg.go.dev/log/slog)
// pentru output structurat.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// Simpla apelare a unor funcții precum `Println` din
	// pachetul `log` folosește logger-ul _standard_, care
	// este deja preconfigurat pentru un output de log
	// rezonabil în `os.Stderr`. Metode suplimentare
	// precum `Fatal*` sau `Panic*` vor opri programul
	// după scrierea în log.
	log.Println("standard logger")

	// Logger-ele pot fi configurate cu _flag-uri_ care
	// le stabilesc formatul output-ului. Implicit,
	// logger-ul standard are setate flag-urile
	// `log.Ldate` și `log.Ltime`, reunite în
	// `log.LstdFlags`. Putem să-i schimbăm flag-urile,
	// de exemplu, ca să afișeze timpul cu precizie de
	// microsecunde.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Suportă și afișarea numelui fișierului și a liniei
	// din care este apelată funcția `log`.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Poate fi util să creăm un logger personalizat și
	// să-l transmitem mai departe. Când creăm un logger
	// nou, putem seta un _prefix_ pentru a-i deosebi
	// output-ul de cel al altor logger-e.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Putem seta prefixul logger-elor existente
	// (inclusiv al celui standard) cu metoda
	// `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Logger-ele pot avea destinații de output
	// personalizate; merge orice `io.Writer`.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Acest apel scrie output-ul de log în `buf`.
	buflog.Println("hello")

	// Iar aici îl afișăm efectiv la ieșirea standard.
	fmt.Print("from buflog:", buf.String())

	// Pachetul `slog` oferă output de log _structurat_.
	// De exemplu, scrierea de loguri în format JSON
	// este simplă.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// Pe lângă mesaj, output-ul `slog` poate conține
	// oricâte perechi cheie=valoare.
	myslog.Info("hello again", "key", "val", "age", 25)
}
