# Încearcă să rulezi codul de scriere a fișierelor.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Apoi verifică conținutul fișierelor scrise.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# În continuare vom vedea cum aplicăm câteva dintre
# ideile de I/O cu fișiere pe care tocmai le-am văzut
# la stream-urile `stdin` și `stdout`.
