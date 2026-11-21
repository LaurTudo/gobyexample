// Go oferă suport integrat pentru
// [codificarea/decodificarea base64](https://en.wikipedia.org/wiki/Base64).

package main

// Această sintaxă importă pachetul `encoding/base64` cu
// numele `b64` în loc de cel implicit, `base64`. Ne va
// economisi puțin spațiu mai jos.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Iată `string`-ul pe care îl vom
	// codifica/decodifica.
	data := "abc123!?$*&()'-=@~"

	// Go suportă atât base64 standard, cât și base64
	// compatibil cu URL-urile. Iată cum codificăm folosind
	// encoder-ul standard. Encoder-ul are nevoie de un
	// `[]byte`, așa că ne convertim `string`-ul la acest
	// tip.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decodificarea poate returna o eroare, pe care o poți
	// verifica dacă nu știi deja că input-ul este bine
	// format.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Aici codificăm/decodificăm folosind un format
	// base64 compatibil cu URL-urile.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
