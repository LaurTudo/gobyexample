// Scrierea unui server HTTP simplu este ușoară cu
// pachetul `net/http`.
package main

import (
	"fmt"
	"net/http"
)

// Un concept fundamental al serverelor `net/http` îl
// reprezintă *handler-ele*. Un handler este un obiect
// care implementează interfața `http.Handler`. Un mod
// obișnuit de a scrie un handler este folosirea
// adaptorului `http.HandlerFunc` pe funcții cu
// semnătura potrivită.
func hello(w http.ResponseWriter, req *http.Request) {

	// Funcțiile care servesc drept handler-e primesc ca
	// argumente un `http.ResponseWriter` și un
	// `http.Request`. Response writer-ul este folosit
	// pentru a completa răspunsul HTTP. Aici, răspunsul
	// nostru simplu este doar "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Acest handler face ceva puțin mai sofisticat:
	// citește toate header-ele cererii HTTP și le
	// reproduce în corpul răspunsului.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Înregistrăm handler-ele pe rutele serverului
	// folosind funcția comodă `http.HandleFunc`. Aceasta
	// configurează *router-ul implicit* din pachetul
	// `net/http` și primește o funcție ca argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// În final, apelăm `ListenAndServe` cu portul și un
	// handler. `nil` îi spune să folosească router-ul
	// implicit pe care tocmai l-am configurat.
	http.ListenAndServe(":8090", nil)
}
