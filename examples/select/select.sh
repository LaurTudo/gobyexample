# Primim valorile `"one"` și apoi `"two"`, cum era de
# așteptat.
$ time go run select.go 
received one
received two

# Observă că timpul total de execuție este doar de ~2
# secunde, pentru că ambele apeluri `Sleep`, de 1 și de
# 2 secunde, se execută concurent.
real	0m2.245s
