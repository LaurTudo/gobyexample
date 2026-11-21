# Observă că, deși slice-urile sunt tipuri diferite de
# array-uri, `fmt.Println` le afișează în mod similar.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Citește acest
# [articol excelent](https://go.dev/blog/slices-intro)
# scris de echipa Go pentru mai multe detalii despre
# designul și implementarea slice-urilor în Go.

# Acum că am văzut array-urile și slice-urile, vom
# trece la cealaltă structură de date predefinită
# esențială din Go: map-urile.
