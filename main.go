package main

type Location struct {
	x int
	y int
}

type Attacker struct {
	name string
	hp   int
	atk  int
	def  int
	loca Location
}

func NewAttacker(name string) *Attacker {

	return &Attacker{name: name, hp: 100, atk: 10, def: 8, loca: Location{x: 10, y: 10}}
}

func (attacker *Attacker) getAttack(oppAttacker Attacker) {
	attacker.hp -= oppAttacker.atk - attacker.def
	println(oppAttacker.name, "attacks", attacker.name)
}

func main() {
	attackLoc := Location{x: 10, y: 10}
	attacker1 := Attacker{name: "attacker1", hp: 100, atk: 10, def: 8, loca: attackLoc}
	attacker2 := Attacker{name: "attacker2", hp: 100, atk: 10, def: 8, loca: attackLoc}
	attacker3 := NewAttacker("attacker3")
	attacker1.getAttack(attacker2)

	println(attacker1.hp, attacker2.hp, attacker3.def)
}
