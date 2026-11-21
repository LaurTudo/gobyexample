# Folosește aceste comenzi pentru a rula exemplul.
# (Notă: din cauza limitărilor Go Playground, acest
# exemplu poate fi rulat doar pe mașina ta locală.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

