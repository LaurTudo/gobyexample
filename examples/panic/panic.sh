# Rularea acestui program îl va face să intre în panică,
# să afișeze un mesaj de eroare și trace-urile
# goroutine-urilor și să se termine cu un cod de ieșire
# diferit de zero.

# Când se declanșează primul panic din `main`,
# programul se oprește fără să ajungă la restul
# codului. Dacă vrei să vezi programul încercând să
# creeze un fișier temporar, comentează primul panic.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Observă că, spre deosebire de unele limbaje care
# folosesc excepții pentru tratarea multor erori, în Go
# este idiomatic să folosești valori returnate care
# indică erorile, oriunde este posibil.
