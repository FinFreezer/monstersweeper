package monstersweeper

type Player struct {
	Health       int
	Strength     int
	Dexterity    int
	Intelligence int
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
