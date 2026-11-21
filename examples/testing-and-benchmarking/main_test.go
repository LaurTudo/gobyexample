// Testarea unitară (unit testing) este o parte
// importantă a scrierii unor programe Go bine
// structurate. Pachetul `testing` oferă instrumentele
// de care avem nevoie pentru a scrie teste unitare, iar
// comanda `go test` rulează testele.

// Pentru demonstrație, acest cod se află în pachetul
// `main`, dar ar putea fi în orice pachet. Codul de test
// se află de obicei în același pachet cu codul pe care
// îl testează.
package main

import (
	"fmt"
	"testing"
)

// Vom testa această implementare simplă a minimului a
// două numere întregi. De obicei, codul pe care îl
// testăm s-ar afla într-un fișier sursă cu un nume de
// genul `intutils.go`, iar fișierul de test pentru el
// s-ar numi `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Un test se creează scriind o funcție al cărei nume
// începe cu `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` raportează eșecurile testului, dar
		// continuă execuția lui. `t.Fatal*` raportează
		// eșecurile și oprește testul imediat.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Scrierea testelor poate fi repetitivă, așa că este
// idiomatic să folosim un *stil bazat pe tabele*
// (table-driven), în care input-urile testelor și
// output-urile așteptate sunt listate într-un tabel,
// iar o singură buclă le parcurge și execută logica de
// test.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// `t.Run` permite rularea de „subteste”, câte unul
		// pentru fiecare intrare din tabel. Acestea sunt
		// afișate separat la rularea `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Testele de benchmark se pun de obicei în fișiere
// `_test.go` și au nume care încep cu `Benchmark`.
// Orice cod necesar pentru rularea benchmark-ului, dar
// care nu trebuie măsurat, se pune înaintea acestei
// bucle.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// Runner-ul de benchmark va executa automat corpul
		// acestei bucle de multe ori, pentru a obține o
		// estimare rezonabilă a timpului de execuție al
		// unei singure iterații.
		IntMin(1, 2)
	}
}
