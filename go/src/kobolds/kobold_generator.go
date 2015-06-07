package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
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
	brawn := rollDice(2)
	return Kobold{"klak", brawn, rollDice(2), rollDice(2), rollDice(2), brawn,
		4, 3, 2, 1, p, p, p, armor, gear, weapon}
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

func init() {
	http.HandleFunc("/", handler)
}

func rollDice(num_dice int) int {
	result := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num_dice; i++ {

		result += r.Intn(7)
	}
	return result
}
