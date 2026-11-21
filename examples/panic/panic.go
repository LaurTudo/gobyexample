// Un `panic` înseamnă de obicei că ceva a mers prost
// în mod neașteptat. În general îl folosim pentru a
// eșua rapid la erori care n-ar trebui să apară în
// funcționarea normală sau pe care nu suntem pregătiți
// să le tratăm elegant.

package main

import (
	"os"
	"path/filepath"
)

func main() {

	// Vom folosi panic pe tot site-ul pentru a verifica
	// erorile neașteptate. Acesta este singurul program
	// de pe site conceput să intre în panică.
	panic("a problem")

	// O utilizare obișnuită a lui panic este oprirea
	// programului dacă o funcție returnează o eroare pe
	// care nu știm (sau nu vrem) s-o tratăm. Iată un
	// exemplu de `panic` în cazul în care primim o
	// eroare neașteptată la crearea unui fișier nou.
	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}
