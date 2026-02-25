package monstersweeper

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Monster struct {
	Name          string
	MaxHealth     int
	Health        int
	Strength      int
	Dexterity     int
	Intelligence  int
	MonsterId     int
	KeyCarrier    bool
	Abilities     map[string]int //Name maps to the amount carried.
	AnimFrames    map[string][]*ebiten.Image
	PrevFrameTime time.Time
	CurrentFrame  int
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
	case 1:
		m = newImp(id)
		return m
	case 2:
		m = newSkeleton(id)
		return m
	case 3:
		m = newZombie(id)
		return m
	case 4:
		m = newWitch(id)
		return m
	case 5:
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
				baseDodgeChance := (target.getDexterity() % 8) / 2

				switch t := target.(type) {
				case *Player:
					if t.PrimaryStat == "Dexterity" {
						if rollAgainstDex(baseDodgeChance) {
							return true
						} else {
							continue
						}
					} else {
						if rollWithoutDex(baseDodgeChance) {
							return true
						} else {
							continue
						}
					}

				case *Monster:
					if rollWithoutDex(baseDodgeChance) {
						return true
					} else {
						continue
					}

				default:
					return false
				}
			} else {
				return true
			}
		} else {
			continue
		}
	}
	return false
}

func rollAgainstDex(baseDodgeChance int) bool {
	rollHitDodge := RollDice(10, 2)
	if rollHitDodge[0] > baseDodgeChance && rollHitDodge[1] > baseDodgeChance {
		return true
	} else {
		return false
	}
}

func rollWithoutDex(baseDodgeChance int) bool {
	rollHitDodge := RollDice(10, 1)
	if rollHitDodge[0] > baseDodgeChance {
		return true
	} else {
		return false
	}
}
func (m *Monster) dealDamage(target Actor) {
	damageRoll := RollDice(2, 2)
	damage := m.Strength/3 + damageRoll[0] + damageRoll[1]
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

func (m *Monster) getHealth() int {
	return m.Health
}

func (m *Monster) getName() string {
	return m.Name
}
