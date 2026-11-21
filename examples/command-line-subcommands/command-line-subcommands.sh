$ go build command-line-subcommands.go 

# Mai întâi apelăm subcomanda foo.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Acum încercăm bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Dar bar nu acceptă flag-urile lui foo.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# În continuare vom vedea variabilele de mediu, un alt
# mod obișnuit de a parametriza programele.
