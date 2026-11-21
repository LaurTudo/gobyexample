// Go oferă suport integrat pentru crearea de conținut
// dinamic sau pentru afișarea unui output personalizat
// către utilizator, prin pachetul `text/template`. Un
// pachet înrudit, numit `html/template`, oferă același
// API, dar are funcționalități de securitate în plus și
// ar trebui folosit pentru generarea de HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// Putem crea un template nou și îi putem parsa
	// conținutul dintr-un string.
	// Template-urile sunt un amestec de text static și
	// „acțiuni” încadrate în `{{...}}`, folosite pentru
	// a insera conținut în mod dinamic.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Alternativ, putem folosi funcția `template.Must`
	// pentru a intra în panică în cazul în care `Parse`
	// returnează o eroare. Asta e util mai ales pentru
	// template-urile inițializate în scope-ul global.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// „Executând” template-ul, îi generăm textul cu
	// valori concrete pentru acțiunile lui. Acțiunea
	// `{{.}}` este înlocuită cu valoarea transmisă ca
	// parametru către `Execute`.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Funcție ajutătoare pe care o vom folosi mai jos.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Dacă datele sunt un struct, putem folosi acțiunea
	// `{{.FieldName}}` pentru a-i accesa câmpurile.
	// Câmpurile trebuie să fie exportate pentru a fi
	// accesibile în timpul execuției unui template.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Același lucru este valabil pentru map-uri; la
	// map-uri nu există nicio restricție privind
	// literele mari sau mici din numele cheilor.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else oferă execuție condiționată în template-uri.
	// O valoare este considerată falsă dacă este valoarea
	// implicită a unui tip, cum ar fi 0, un string gol,
	// un pointer nil etc.
	// Acest exemplu arată și o altă funcționalitate a
	// template-urilor: folosirea lui `-` în acțiuni
	// pentru a elimina spațiile albe.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// Blocurile range ne permit să parcurgem slice-uri,
	// array-uri, map-uri sau canale. În interiorul
	// blocului range, `{{.}}` primește elementul curent
	// al iterației.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
