// Uneori vrem ca programele noastre Go să trateze
// inteligent [semnalele Unix](https://en.wikipedia.org/wiki/Unix_signal).
// De exemplu, am putea vrea ca un server să se oprească
// elegant (graceful shutdown) când primește un
// `SIGTERM`, sau ca o unealtă din linia de comandă să
// oprească procesarea input-ului dacă primește un
// `SIGINT`. Iată un mod modern de a trata semnalele
// folosind contexte.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// `signal.NotifyContext` returnează un context care
	// este anulat când sosește unul dintre semnalele
	// enumerate.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Programul va aștepta aici până când primește unul
	// dintre semnalele configurate.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// `context.Cause` raportează de ce a fost anulat
	// contextul. Pentru o anulare declanșată de un semnal,
	// aceasta include valoarea semnalului.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
