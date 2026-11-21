// URL-urile oferă un [mod uniform de a localiza resurse](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Iată cum se parsează URL-urile în Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Vom parsa acest URL de exemplu, care include o
	// schemă, informații de autentificare, host, port,
	// cale (path), parametri de query și fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parsăm URL-ul și ne asigurăm că nu există erori.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accesarea schemei este simplă.
	fmt.Println(u.Scheme)

	// `User` conține toate informațiile de autentificare;
	// apelează `Username` și `Password` pe el pentru
	// valorile individuale.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` conține atât numele host-ului, cât și
	// portul, dacă există. Folosește `SplitHostPort`
	// pentru a le extrage.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Aici extragem `path`-ul și fragmentul de după
	// `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Pentru a obține parametrii de query ca string în
	// formatul `k=v`, folosește `RawQuery`. Poți parsa
	// parametrii de query și într-un map. Map-urile
	// rezultate sunt de la string-uri la slice-uri de
	// string-uri, așa că folosește indexul `[0]` dacă
	// vrei doar prima valoare.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
