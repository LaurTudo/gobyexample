# Când rulăm programul, mesajul `"ping"` este transmis
# cu succes dintr-o goroutine în alta prin canalul
# nostru.
$ go run channels.go 
ping

# Implicit, trimiterile și primirile se blochează până
# când atât emițătorul, cât și receptorul sunt pregătiți.
# Această proprietate ne-a permis să așteptăm mesajul
# `"ping"` la finalul programului fără să folosim
# nicio altă sincronizare.
