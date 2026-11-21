// Parsarea numerelor din string-uri este o operație
// simplă, dar frecventă în multe programe; iată cum se
// face în Go.

package main

// Pachetul predefinit `strconv` se ocupă de parsarea
// numerelor.
import (
	"fmt"
	"strconv"
)

func main() {

	// La `ParseFloat`, acest `64` indică pe câți biți de
	// precizie se face parsarea.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// La `ParseInt`, `0` înseamnă că baza se deduce din
	// string. `64` cere ca rezultatul să încapă în 64 de
	// biți.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` recunoaște numerele în format
	// hexazecimal.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Este disponibilă și funcția `ParseUint`.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` este o funcție comodă pentru parsarea simplă
	// a unui `int` în baza 10.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Funcțiile de parsare returnează o eroare la un
	// input invalid.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
