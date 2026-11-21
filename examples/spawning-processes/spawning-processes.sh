# Programele lansate returnează același output ca și
# cum le-am fi rulat direct din linia de comandă.
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# date nu are flag-ul `-x`, așa că se va termina cu un
# mesaj de eroare și un cod de ieșire diferit de zero.
command exit rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
