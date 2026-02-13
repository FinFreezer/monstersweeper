package monstersweeper

import (
	"log"
)

type Monster struct {
	Name         string
	Health       int
	Strength     int
	Dexterity    int
	Intelligence int
	MonsterId    int
	KeyCarrier   bool
	Abilities    map[string]int //Name maps to the amount carried.
}

type MonsterId int

const (
	Imp MonsterId = iota
	Skeleton
	Zombie
	Witch
	Orc
)

var (
	monsters = []string{"Imp", "Skeleton", "Zombie", "Witch", "Orc"}
)

func InitMonster() {
	return
}

func newMonster(id int) (m *Monster) {
	switch m.MonsterId {
	case 0:
		newImp(m.MonsterId)
	case 1:
		newSkeleton(m.MonsterId)
	case 2:
		newZombie(m.MonsterId)
	case 3:
		newWitch(m.MonsterId)
	case 4:
		newOrc(m.MonsterId)
	default:
		log.Fatal("Unable to reach monster.")
	}
	return
}

func newImp(id int) *Monster {
	m := Monster{
		Name:         "Imp",
		Health:       5,
		Strength:     6,
		Dexterity:    8,
		Intelligence: 8,
		MonsterId:    id,
	}
	return &m
}

func newSkeleton(id int) *Monster {
	m := Monster{
		Name:         "Skeleton",
		Health:       8,
		Strength:     10,
		Dexterity:    5,
		Intelligence: 5,
		MonsterId:    id,
	}
	return &m
}

func newZombie(id int) *Monster {
	m := Monster{
		Name:         "Zombie",
		Health:       12,
		Strength:     8,
		Dexterity:    4,
		Intelligence: 2,
		MonsterId:    id,
	}
	return &m
}

func newWitch(id int) *Monster {
	m := Monster{
		Name:         "Witch",
		Health:       4,
		Strength:     6,
		Dexterity:    10,
		Intelligence: 10,
		MonsterId:    id,
	}
	return &m
}

func newOrc(id int) *Monster {
	m := Monster{
		Name:         "Orc",
		Health:       14,
		Strength:     10,
		Dexterity:    4,
		Intelligence: 2,
		MonsterId:    id,
	}
	return &m
}
