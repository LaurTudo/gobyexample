# Rulăm serverul în fundal.
$ go run context.go &

# Simulăm o cerere a unui client către `/hello` și
# apăsăm Ctrl+C la scurt timp după pornire, pentru a
# semnala anularea.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
