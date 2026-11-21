# Ne așteptăm să obținem exact 50.000 de operații. Dacă
# am fi folosit un întreg non-atomic și l-am fi
# incrementat cu `ops++`, probabil am fi obținut un alt
# număr, diferit de la o rulare la alta, pentru că
# goroutine-urile s-ar fi încurcat unele pe altele. Mai
# mult, am fi primit erori de tip data race la rularea
# cu flag-ul `-race`.
$ go run atomic-counters.go
ops: 50000

# În continuare vom vedea mutex-urile, un alt instrument
# pentru gestionarea stării.
