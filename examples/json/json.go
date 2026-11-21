// Go oferă suport integrat pentru codificarea și
// decodificarea JSON, inclusiv din și în tipuri de date
// predefinite și personalizate.

package main

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Vom folosi aceste două struct-uri pentru a demonstra
// mai jos codificarea și decodificarea tipurilor
// personalizate.
type response1 struct {
	Page   int
	Fruits []string
}

// Doar câmpurile exportate vor fi codificate/decodificate
// în JSON. Câmpurile trebuie să înceapă cu literă mare
// pentru a fi exportate.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Mai întâi vom vedea codificarea tipurilor de date
	// de bază în string-uri JSON. Iată câteva exemple
	// pentru valori atomice.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Iar aici câteva pentru slice-uri și map-uri, care
	// se codifică, cum era de așteptat, în array-uri și
	// obiecte JSON.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Pachetul JSON poate codifica automat tipurile tale
	// de date personalizate. Va include în output-ul
	// codificat doar câmpurile exportate și va folosi
	// implicit numele lor drept chei JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Poți folosi tag-uri în declarațiile câmpurilor
	// struct-ului pentru a personaliza numele cheilor
	// JSON codificate. Uită-te la definiția lui
	// `response2` de mai sus pentru un exemplu de astfel
	// de tag-uri.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Acum să vedem decodificarea datelor JSON în valori
	// Go. Iată un exemplu pentru o structură de date
	// generică.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Trebuie să oferim o variabilă în care pachetul
	// JSON să pună datele decodificate. Acest
	// `map[string]any` va conține un map de la string-uri
	// la tipuri de date arbitrare.
	var dat map[string]any

	// Iată decodificarea propriu-zisă și o verificare a
	// erorilor asociate.
	// Pentru concizie, ignorăm erorile în aceste
	// exemple; în codul real, ar trebui să verifici
	// mereu erorile și să acționezi în consecință.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Pentru a folosi valorile din map-ul decodificat,
	// va trebui să le convertim la tipul potrivit. De
	// exemplu, aici convertim valoarea din `num` la tipul
	// așteptat, `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accesarea datelor imbricate necesită o serie de
	// conversii.
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Putem decodifica JSON și în tipuri de date
	// personalizate. Avantajele sunt că programele
	// noastre câștigă un plus de siguranță a tipurilor
	// și că dispare nevoia de type assertions la
	// accesarea datelor decodificate.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// În exemplele de mai sus am folosit mereu bytes și
	// string-uri ca intermediari între date și
	// reprezentarea JSON afișată la ieșirea standard.
	// Putem trimite codificările JSON și direct, ca
	// stream, către `io.Writer`-e precum `os.Stdout` sau
	// chiar corpurile răspunsurilor HTTP.
	d := map[string]int{"apple": 5, "lettuce": 7}
	var buf bytes.Buffer
	_ = json.MarshalWrite(&buf, d)
	fmt.Println(buf.String())

	// Citirea ca stream din `io.Reader`-e precum
	// `os.Stdin` sau corpurile cererilor HTTP se face cu
	// `json.UnmarshalRead`.
	res1 := response2{}
	_ = json.UnmarshalRead(strings.NewReader(str), &res1)
	fmt.Println(res1)
}
