# Pentru a încerca filtrul de linii, creează mai întâi
# un fișier cu câteva linii scrise cu litere mici.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Apoi folosește filtrul de linii pentru a obține
# liniile cu majuscule.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
