# Dacă rulezi `exit.go` cu `go run`, ieșirea va fi
# preluată de `go` și afișată.
$ go run exit.go
exit status 3

# Compilând și executând un binar, poți vedea
# statusul în terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Observă că `!` din programul nostru nu a fost afișat
# niciodată.
