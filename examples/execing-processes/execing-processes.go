// În exemplul anterior am văzut
// [lansarea de procese externe](spawning-processes).
// Facem asta când avem nevoie ca un proces extern să fie
// accesibil unui proces Go care rulează. Uneori însă
// vrem pur și simplu să înlocuim complet procesul Go
// curent cu un altul (poate unul care nu e scris în Go).
// Pentru asta vom folosi implementarea din Go a
// clasicei funcții
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// În exemplul nostru vom face exec pe `ls`. Go cere o
	// cale absolută către binarul pe care vrem să-l
	// executăm, așa că vom folosi `exec.LookPath` pentru
	// a-l găsi (probabil `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` cere argumentele sub formă de slice (nu ca
	// un singur string mare). Îi vom da lui `ls` câteva
	// argumente obișnuite. Observă că primul argument
	// trebuie să fie numele programului.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` are nevoie și de un set de
	// [variabile de mediu](environment-variables) pe care
	// să le folosească. Aici îi dăm pur și simplu mediul
	// nostru curent.
	env := os.Environ()

	// Iată apelul propriu-zis `syscall.Exec`. Dacă acest
	// apel reușește, execuția procesului nostru se va
	// încheia aici și va fi înlocuită de procesul
	// `/bin/ls -a -l -h`. Dacă apare o eroare, vom primi
	// o valoare returnată.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
