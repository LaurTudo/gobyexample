// În exemplul anterior am văzut cum se configurează un
// [server HTTP](http-server) simplu. Serverele HTTP sunt
// utile pentru a demonstra folosirea lui
// `context.Context` pentru controlul anulării. Un
// `Context` transportă termene limită (deadlines),
// semnale de anulare și alte valori specifice unei
// cereri, peste granițele API-urilor și între
// goroutine-uri.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Mecanismul din `net/http` creează un
	// `context.Context` pentru fiecare cerere, disponibil
	// prin metoda `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Așteptăm câteva secunde înainte de a trimite un
	// răspuns clientului. Asta ar putea simula o muncă pe
	// care o face serverul. În timp ce lucrăm, urmărim
	// canalul `Done()` al contextului pentru un semnal că
	// ar trebui să anulăm munca și să ne întoarcem cât
	// mai repede posibil.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Metoda `Err()` a contextului returnează o eroare
		// care explică de ce a fost închis canalul
		// `Done()`.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Ca înainte, înregistrăm handler-ul pe ruta "/hello"
	// și pornim serverul.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
