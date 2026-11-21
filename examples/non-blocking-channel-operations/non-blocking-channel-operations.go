// Trimiterile și primirile simple pe canale sunt
// blocante. Totuși, putem folosi `select` cu o ramură
// `default` pentru a implementa trimiteri și primiri
// _neblocante_ și chiar instrucțiuni `select` neblocante
// cu mai multe ramuri.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Iată o primire neblocantă. Dacă pe `messages` este
	// disponibilă o valoare, `select` va intra pe ramura
	// `case` `<-messages` cu acea valoare. Dacă nu, va
	// intra imediat pe ramura `default`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// O trimitere neblocantă funcționează similar. Aici
	// `msg` nu poate fi trimis în canalul `messages`,
	// pentru că acesta nu are buffer și nu există niciun
	// receptor. Prin urmare, este aleasă ramura `default`.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Putem folosi mai multe ramuri `case` înaintea
	// ramurii `default` pentru a implementa un select
	// neblocant cu mai multe ramuri. Aici încercăm
	// primiri neblocante atât pe `messages`, cât și pe
	// `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
