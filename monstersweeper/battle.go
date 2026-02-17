package monstersweeper

import (
	"fmt"
	"time"
)

func StartBattle(p *Player, m *Monster) string {
	var debug string
	for {
		time.Sleep(750 * time.Millisecond)
		fmt.Println("Player turn...")
		if p.rollAccuracy(m) {
			hp := m.getHealth()
			p.dealDamage(m)
			debug += fmt.Sprintf("Player dealt %d damage. \n", hp-m.getHealth())
			debug += fmt.Sprintf("%s health remaining: %d \n", m.getName(), m.getHealth())
			if m.isDead() {
				if m.KeyCarrier {
					if p.Items == nil {
						p.Items = make(map[string]int)
					}
					p.Items["Key"] = 1
					fmt.Println("You found the key!")
				}
				fmt.Println("Battle finished. Monster defeated.")
				healChance := RollDice(3, 1)
				debug += fmt.Sprintf("Player won with %d health left. \n", p.getHealth())
				if healChance[0] == 2 {
					fmt.Println("Looted a healing a potion! Healing to full.")
					debug += fmt.Sprintf("Healed. \n")
					p.Health = p.MaxHealth
				}
				debug += "\n\n"
				return debug
			}
		} else {
			fmt.Println("Attack missed.")
		}
		time.Sleep(750 * time.Millisecond)
		fmt.Println("Monster turn...")
		if m.rollAccuracy(p) {
			hp := p.getHealth()
			m.dealDamage(p)
			debug += fmt.Sprintf("%s dealt %d damage. \n", m.getName(), hp-p.getHealth())
			debug += fmt.Sprintf("Player health remaining: %d \n", p.getHealth())
			if p.isDead() {
				debug += fmt.Sprintf("%s won with %d health left. \n", m.getName(), m.getHealth())
				fmt.Println("Battle finished. You died.")
				debug += "\n\n"
				return debug
			}
		} else {
			fmt.Println("Attack missed.")
		}
	}
}
