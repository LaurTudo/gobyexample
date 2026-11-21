# Rulând programul, vedem că contoarele au fost
# actualizate cum ne așteptam.
$ go run mutexes.go
map[a:20000 b:10000]

# În continuare vom vedea cum putem implementa aceeași
# gestionare a stării folosind doar goroutine-uri și
# canale.
