package monstersweeper

type Player struct {
	Name         string
	MaxHealth    int
	Health       int
	Strength     int
	Dexterity    int
	Intelligence int
	PrimaryStat  string
	Items        map[string]int //Name maps to the amount carried.
}

func (p *Player) HasKey() bool {
	if val, ok := p.Items["Key"]; ok {
		if val > 0 {
			return true
		}
	}
	return false
}

func NewPlayer() *Player {
	p := Player{
		Name:         "Knight",
		MaxHealth:    30,
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
			} else {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func (p *Player) dealDamage(target Actor) {
	rollDamage := RollDice(3, 2)
	damage := p.Strength/3 + rollDamage[0] + rollDamage[1]
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

func (p *Player) getHealth() int {
	return p.Health
}

func (p *Player) getName() string {
	return p.Name
}
