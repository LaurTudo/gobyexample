# Rulând programul, vedem că exemplul de gestionare a
# stării bazat pe goroutine-uri realizează în total
# aproximativ 80.000 de operații.
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177

# În acest caz particular, abordarea bazată pe
# goroutine-uri a fost puțin mai complicată decât cea
# bazată pe mutex-uri. Totuși, poate fi utilă în anumite
# situații, de exemplu când sunt implicate și alte
# canale sau când gestionarea mai multor astfel de
# mutex-uri ar fi predispusă la erori. Folosește
# abordarea care ți se pare cea mai naturală, mai ales
# în ceea ce privește înțelegerea corectitudinii
# programului tău.
