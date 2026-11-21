// Go oferă suport integrat pentru XML și formate
// asemănătoare cu XML prin pachetul `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant va fi mapat în XML. La fel ca în exemplele
// JSON, tag-urile câmpurilor conțin directive pentru
// encoder și decoder. Aici folosim câteva
// funcționalități speciale ale pachetului XML: numele de
// câmp `XMLName` stabilește numele elementului XML care
// reprezintă acest struct; `id,attr` înseamnă că
// câmpul `Id` este un _atribut_ XML, nu un element
// imbricat.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Generăm XML-ul care reprezintă planta noastră;
	// folosim `MarshalIndent` pentru a produce un
	// output mai ușor de citit de oameni.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Pentru a adăuga un header XML generic la output,
	// îl adăugăm explicit.
	fmt.Println(xml.Header + string(out))

	// Folosim `Unmarshal` pentru a parsa un stream de
	// bytes cu XML într-o structură de date. Dacă XML-ul
	// este malformat sau nu poate fi mapat pe Plant, va
	// fi returnată o eroare descriptivă.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// Tag-ul de câmp `parent>child>plant` îi spune
	// encoder-ului să imbrice toate elementele `plant`
	// sub `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
