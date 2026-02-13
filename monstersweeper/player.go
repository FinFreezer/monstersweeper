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

func (p *Player) rollDie() {
	return
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
