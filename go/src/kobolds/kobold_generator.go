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
	gear := getGear()
	armor := Armor{"ring", "nil", "nil", 3}
	brawn := rollDice(2)
	meat := calculateHandyNumber(brawn)
	ego := rollDice(2)
	cunning := calculateHandyNumber(ego)
	extraneous := rollDice(2)
	luck := calculateHandyNumber(extraneous)
	reflexes := rollDice(2)
	agility := calculateHandyNumber(reflexes)
	return Kobold{"klak", brawn, ego, extraneous, reflexes, brawn,
		meat, cunning, luck, agility, p, p, p, armor, gear, weapon}
}

func getGear() Gear {
	// TODO(whillegas): make this a constant so we don't invoke it each time.
	gearMap := map[int]Gear{
		2:  Gear{"Cup of milk elemental summoning", "", 0, 0},
		3:  Gear{"Bag of holding: chickens", "", 0, 0},
		4:  Gear{"Ring of human speaking", "", 0, 0},
		5:  Gear{"Codex of Tabriz the Arcane", "", 0, 0},
		6:  Gear{"Random Booze", "", 0, 0},
		7:  Gear{"Spice Sack", "", 0, 0},
		8:  Gear{"BackPack", "", 0, 0},
		9:  Gear{"Adventurer's cast-offs", "", 0, 0},
		10: Gear{"25 feet of rope", "", 0, 0},
		11: Gear{"10 Foot Pole", "", 0, 0},
		12: Gear{"lint", "", 0, 0},
	}

	return gearMap[rollDice(2)]

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

func calculateHandyNumber(stat int) int {

	if stat < 5 {
		return 1
	}

	if stat < 9 {
		return 2
	}

	if stat < 13 {
		return 3
	}

	if stat < 17 {
		return 4
	}

	if stat < 21 {
		return 5
	}
	return 0
}
