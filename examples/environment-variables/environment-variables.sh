# Rulând programul, vedem că preluăm valoarea lui `FOO`
# pe care am setat-o în program, dar că `BAR` este
# gol.
$ go run environment-variables.go
FOO: 1
BAR: 

# Lista cheilor din mediu va depinde de mașina ta.
TERM_PROGRAM
PATH
SHELL
...
FOO

# Dacă setăm mai întâi `BAR` în mediu, programul rulat
# preia acea valoare.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
