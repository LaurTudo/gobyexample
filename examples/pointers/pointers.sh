# `zeroval` nu modifică variabila `i` din `main`, dar
# `zeroptr` o modifică, pentru că are o referință la
# adresa de memorie a acelei variabile.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
value at *p: 42
value at *p: 0