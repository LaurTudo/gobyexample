// Pachetul `net` oferă uneltele de care avem nevoie
// pentru a construi ușor servere cu socket-uri TCP.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// `net.Listen` pornește serverul pe rețeaua dată
	// (TCP) și pe adresa dată (portul 8090 pe toate
	// interfețele).
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// Închidem listener-ul pentru a elibera portul
	// când aplicația se termină.
	defer listener.Close()

	// Buclă infinită care acceptă conexiuni noi de la
	// clienți.
	for {
		// Așteptăm o conexiune.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// Folosim aici o goroutine pentru a trata
		// conexiunea, astfel încât bucla principală să
		// poată accepta în continuare alte conexiuni.
		go handleConnection(conn)
	}
}

// `handleConnection` tratează o singură conexiune cu un
// client: citește o linie de text de la client și
// returnează un răspuns.
func handleConnection(conn net.Conn) {
	// Închiderea conexiunii eliberează resursele când am
	// terminat interacțiunea cu clientul.
	defer conn.Close()

	// Folosim `bufio.NewReader` pentru a citi o linie de
	// date de la client (terminată cu un newline).
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// Creăm și trimitem un răspuns înapoi clientului,
	// demonstrând comunicarea în ambele sensuri.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
