# Rularea programului calculează hash-ul și îl afișează
# într-un format hexazecimal ușor de citit.
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Poți calcula alte hash-uri folosind un tipar similar
# celui de mai sus. De exemplu, pentru a calcula
# hash-uri SHA512, importă `crypto/sha512` și folosește
# `sha512.New()`.

# Observă că, dacă ai nevoie de hash-uri sigure din punct
# de vedere criptografic, ar trebui să te documentezi
# atent despre
# [puterea hash-urilor](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
