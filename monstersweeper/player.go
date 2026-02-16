package monstersweeper

type Player struct {
	Health       int
	Strength     int
	Dexterity    int
	Intelligence int
	PrimaryStat  string
	Items        map[string]int //Name maps to the amount carried.
}

func (p *Player) hasKey() bool {
	if val, ok := p.Items["Key"]; ok {
		if val > 0 {
			return true
		}
	}
	return false
}

func NewPlayer() *Player {
	p := Player{
		Health:       30,
		Strength:     14,
		Dexterity:    10,
		Intelligence: 6,
		PrimaryStat:  "Dexterity",
	}
	return &p
}

func (p *Player) rollAccuracy(target Actor) bool {
	rolls := RollDice(6, 2)

	for _, roll := range rolls {
		if roll >= 4 {
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

func (p *Player) dealDamage(target Actor) {
	rollDamage := RollDice(4, 3)
	damage := p.Strength + rollDamage[0] + rollDamage[1] + rollDamage[2]
	target.takeDamage(damage)
}

func (p *Player) takeDamage(damage int) {
	p.Health -= damage
}

func (p *Player) getDexterity() int {
	return p.Dexterity
}

func (p *Player) isDead() bool {
	if p.Health <= 0 {
		return true
	}
	return false
}
