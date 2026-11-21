// [_Hash-urile SHA256_](https://en.wikipedia.org/wiki/SHA-2)
// sunt folosite frecvent pentru a calcula identificatori
// scurți pentru blob-uri binare sau de text. De exemplu,
// certificatele TLS/SSL folosesc SHA256 pentru a calcula
// semnătura unui certificat. Iată cum se calculează
// hash-uri SHA256 în Go.

package main

// Go implementează mai multe funcții hash în diverse
// pachete `crypto/*`.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Aici pornim cu un hash nou.
	h := sha256.New()

	// `Write` așteaptă bytes. Dacă ai un string `s`,
	// folosește `[]byte(s)` pentru a-l converti în bytes.
	h.Write([]byte(s))

	// Aici obținem rezultatul final al hash-ului ca slice
	// de bytes. Argumentul lui `Sum` poate fi folosit
	// pentru a adăuga la un slice de bytes existent: de
	// obicei nu e necesar.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
