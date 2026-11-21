# Pentru a experimenta cu programul de flag-uri, e cel
# mai bine să-l compilezi mai întâi și apoi să rulezi
# direct binarul rezultat.
$ go build command-line-flags.go

# Încearcă programul compilat, dându-i mai întâi valori
# pentru toate flag-urile.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Observă că, dacă omiți flag-uri, acestea iau automat
# valorile implicite.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Argumentele poziționale pot fi date la final, după
# flag-uri.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Observă că pachetul `flag` cere ca toate flag-urile
# să apară înaintea argumentelor poziționale (altfel,
# flag-urile vor fi interpretate ca argumente
# poziționale).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Folosește flag-urile `-h` sau `--help` pentru a obține
# textul de ajutor generat automat pentru program.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Dacă dai un flag care nu a fost declarat în pachetul
# `flag`, programul va afișa un mesaj de eroare și va
# arăta din nou textul de ajutor.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
