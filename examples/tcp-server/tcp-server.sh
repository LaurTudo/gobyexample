# Rulăm serverul TCP în fundal.
$ go run tcp-server.go &

# Trimitem date și preluăm răspunsul folosind netcat.
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT

