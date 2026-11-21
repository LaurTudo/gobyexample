# Când rulăm acest program, vedem mai întâi output-ul
# apelului blocant, apoi output-ul celor două
# goroutine-uri. Output-ul goroutine-urilor poate fi
# intercalat, pentru că goroutine-urile sunt rulate
# concurent de runtime-ul Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# În continuare vom vedea un complement al
# goroutine-urilor în programele Go concurente:
# canalele.
