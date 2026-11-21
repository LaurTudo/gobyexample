// Unele unelte din linia de comandă, precum unealta
// `go` sau `git`, au multe *subcomenzi*, fiecare cu
// propriul set de flag-uri. De exemplu, `go build` și
// `go get` sunt două subcomenzi diferite ale uneltei
// `go`. Pachetul `flag` ne permite să definim ușor
// subcomenzi simple, cu flag-uri proprii.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Declarăm o subcomandă folosind funcția
	// `NewFlagSet`, apoi definim flag-uri noi, specifice
	// acestei subcomenzi.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Pentru o altă subcomandă putem defini alte flag-uri
	// suportate.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Subcomanda este așteptată ca prim argument al
	// programului.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Verificăm ce subcomandă este apelată.
	switch os.Args[1] {

	// Pentru fiecare subcomandă, îi parsăm flag-urile
	// proprii și avem acces la argumentele poziționale
	// de la final.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
