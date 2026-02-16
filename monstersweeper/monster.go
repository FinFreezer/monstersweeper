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

func NewMonster(id int) (m *Monster) {
	switch id {
	case 0:
		m = newImp(id)
		return m
	case 1:
		m = newSkeleton(id)
		return m
	case 2:
		m = newZombie(id)
		return m
	case 3:
		m = newWitch(id)
		return m
	case 4:
		m = newOrc(id)
		return m
	default:
		log.Fatal("Unable to reach monster.")
	}
	return
}

func (m *Monster) rollAccuracy(target Actor) bool {
	rolls := RollDice(6, 2)

	for _, roll := range rolls {
		if roll > 4 {
			if target.getDexterity() >= 10 {
				dodge := (target.getDexterity() % 8) / 2

				rollDodge := RollDice(10, 1)
				if rollDodge[0] < dodge {
					continue
				} else {
					return true
				}

			}
		}
	}
	return false
}

func (m *Monster) dealDamage(target Actor) {
	damageRoll := RollDice(4, 2)
	damage := m.Strength + damageRoll[0] + damageRoll[1]
	target.takeDamage(damage)
}

func (m *Monster) takeDamage(damage int) {
	m.Health -= damage
}

func (m *Monster) getDexterity() int {
	return m.Dexterity
}

func (m *Monster) isDead() bool {
	if m.Health <= 0 {
		return true
	}
	return false
}
