# Când rulăm programul, acesta este înlocuit de `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Observă că Go nu oferă clasica funcție Unix `fork`.
# De obicei însă asta nu e o problemă, pentru că
# pornirea de goroutine-uri, lansarea de procese și
# exec-ul de procese acoperă majoritatea cazurilor de
# utilizare pentru `fork`.
