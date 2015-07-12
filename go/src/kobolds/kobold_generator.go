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
	skills := chooseSkills(ego)
	hits := brawn // This also needs to include modifiers from gear
	return Kobold{"klak", brawn, ego, extraneous, reflexes, hits,
		meat, cunning, luck, agility, skills, p, p, armor, gear, weapon}
}

func chooseSkills(ego int) []string {
	numskills := 6
	if ego < 6 {
		numskills = ego
	}
	skills := make([]string, numskills)

	for i := range skills {
		skills[i] = getSkill(i)

	}

	// Need to generate at least one of each skill type, and then Random
	// for coming back to this.

	return skills
}

func getSkill(i int) string {
	switch i {
	case 0:
		return "Cook"
	case 1:
		return getBrawnSkill()
	case 2:
		return getEgoSkill()
	case 3:
		return getExtraneousSkill()
	case 4:
		return getReflexesSkill()
	default:
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return getSkill(r.Intn(3) + 1)
	}

}

func getBrawnSkill() string {
	brawnSkillMap := map[int]string{
		1: "Athlete",
		2: "Bully",
		3: "Duel",
		4: "Lift",
		5: "Swim",
		6: "Wrassle",
	}
	return brawnSkillMap[rollDice(1)]
}

func getEgoSkill() string {
	skillMap := map[int]string{
		1: "Hide",
		2: "Lackey",
		3: "Sage",
		4: "Shoot",
		5: "Speak Human",
		6: "Trap",
	}
	return skillMap[rollDice(1)]
}

func getExtraneousSkill() string {
	skillMap := map[int]string{
		1:  "Bard",
		2:  "Dungeon",
		3:  "Track",
		4:  "Trade",
		5:  "Bard",
		6:  "Dungeon",
		7:  "Track",
		8:  "Trade",
		9:  "Bard",
		10: "Dungeon",
		11: "Track",
		12: "Trade",
	}
	return skillMap[rollDice(3)]
}

func getReflexesSkill() string {
	skillMap := map[int]string{
		1: "Cower",
		2: "Fast",
		3: "Sneak",
		4: "Steal",
		5: "Ride",
		6: "Wiggle",
	}
	return skillMap[rollDice(1)]
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

func rollDice(numDice int) int {
	result := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numDice; i++ {

		result += r.Intn(6) + 1
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
