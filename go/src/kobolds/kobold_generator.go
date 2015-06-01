package main

import "fmt"

// Indicates the skills of a Kobold
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
	name                               string
	brawn, ego, extraneous, reflexes   int
	hits, meat, cunning, luck, agility int
	skills                             []string
	edges                              []string
	bogies                             []string
	armor                              Armor
	gear                               Gear
	weapon                             Weapon
}

func generateKobold() Kobold {
	p := []string{"a"}
	weapon := Weapon{"knife", 1, "nil", "loser"}
	gear := Gear{"hood", "nil", 1, 3}
	armor := Armor{"ring", "nil", "nil", 3}
	return Kobold{"klak", 1, 2, 3, 4, 5, 4, 3, 2, 1, p, p, p, armor, gear, weapon}
}

func displayKobold(kobold Kobold) {
	fmt.Println(kobold)
}

func main() {
	kobold := generateKobold()
	displayKobold(kobold)
}
