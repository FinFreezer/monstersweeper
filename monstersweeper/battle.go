package monstersweeper

import (
	"fmt"
)

func StartBattle(p *Player, m *Monster) {
	fmt.Println("Player turn...")
	if p.rollAccuracy(m) {
		p.dealDamage(m)
	}
}
