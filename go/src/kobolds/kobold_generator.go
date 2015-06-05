package main

import (
	"net/http"
	"text/template"
)

// Weapon Indicates the skills of a Kobold
type Weapon struct {
	name  string
	dam   int
	skill string
	bogie string
}

type Armor struct {
	name  string
	skill string
	bogie string
	dam   int
}

type Gear struct {
	name  string
	skill string
	hits  int
	dam   int
}

type Kobold struct {
	Name                               string
	Brawn, Ego, Extraneous, Reflexes   int
	Hits, Meat, Cunning, Luck, Agility int
	Skills                             []string
	Edges                              []string
	Bogies                             []string
	Armor                              Armor
	Gear                               Gear
	Weapon                             Weapon
}

func generateKobold() Kobold {
	p := []string{"a"}
	weapon := Weapon{"knife", 1, "nil", "loser"}
	gear := Gear{"hood", "nil", 1, 3}
	armor := Armor{"ring", "nil", "nil", 3}
	return Kobold{"klak", 1, 2, 3, 4, 5, 4, 3, 2, 1, p, p, p, armor, gear, weapon}
}

func handler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, , r.URL.Path[1:])
	t, err := template.ParseFiles("templates/kobold_table.html")
	if err != nil {
		panic(err)
	}
	kobold := generateKobold()
	t.Execute(w, kobold)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	// displayKobold(kobold)
}
