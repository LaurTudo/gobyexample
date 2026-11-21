# Când rulăm acest program, se va bloca așteptând un
# semnal. Apăsând `ctrl-C` (pe care terminalul îl
# afișează ca `^C`), putem trimite un semnal `SIGINT`,
# iar programul va afișa cauza anulării și apoi se va
# termina.
$ go run signals.go
awaiting signal
^C
interrupt signal received
exiting
