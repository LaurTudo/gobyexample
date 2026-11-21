// Biblioteca standard Go vine cu un suport excelent
// pentru clienți și servere HTTP, în pachetul
// `net/http`. În acest exemplu îl vom folosi pentru a
// trimite cereri HTTP simple.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Trimitem o cerere HTTP GET către un server.
	// `http.Get` este o scurtătură comodă în locul creării
	// unui obiect `http.Client` și apelării metodei lui
	// `Get`; folosește obiectul `http.DefaultClient`,
	// care are setări implicite utile.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Afișăm statusul răspunsului HTTP.
	fmt.Println("Response status:", resp.Status)

	// Afișăm primele 5 linii din corpul răspunsului.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
