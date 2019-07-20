package main

import (
	"fmt"

	"github.com/nmcapule/graphrng/entity"
)

func main() {
	trainer := entity.NewEntity("Red", "actor")
	pokemon1 := entity.NewEntity("Pikachu", "actor")
	pokemon2 := entity.NewEntity("Bulbasaur", "actor")
	pokemon3 := entity.NewEntity("Mew", "actor")
	sword := entity.NewEntity("Sword of Time", "equip")

	party := entity.NewEntity("Champions", "party")

	entity.Connect("owner", trainer, party, true)
	entity.Connect("member", pokemon1, party, false)
	entity.Connect("member", pokemon2, party, false)
	entity.Connect("member", pokemon3, party, false)
	entity.Connect("hate", pokemon1, pokemon2, false)
	entity.Connect("love", pokemon3, pokemon2, true)

	entity.Connect("wield", pokemon3, sword, true)

	entity.DebugPrint(trainer)
	entity.DebugPrint(pokemon1)
	entity.DebugPrint(pokemon2)
	entity.DebugPrint(pokemon3)
	entity.DebugPrint(party)

	fmt.Println("----- Disbanding party! -----")

	for _, r := range party.Relationships {
		entity.Disband(r)
	}

	entity.DebugPrint(trainer)
	entity.DebugPrint(pokemon1)
	entity.DebugPrint(pokemon2)
	entity.DebugPrint(pokemon3)
	entity.DebugPrint(party)
}
