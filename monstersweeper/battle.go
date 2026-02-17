package monstersweeper

import (
	"fmt"
	"time"
)

func StartBattle(p *Player, m *Monster) {
	for {
		fmt.Println("Player turn...")
		if p.rollAccuracy(m) {
			hp := m.getHealth()
			p.dealDamage(m)
			fmt.Printf("Player dealt %d damage. \n", hp-m.getHealth())
			fmt.Printf("%s health remaining: %d \n", m.getName(), m.getHealth())
			if m.isDead() {
				fmt.Println("Battle finished. Monster defated.")
				healChance := RollDice(3, 1)
				if healChance[0] == 2 {
					fmt.Println("Looted a healing a potion! Healing to full.")
					p.Health = p.MaxHealth
				}
				return
			}
		} else {
			fmt.Println("Attack missed.")
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Monster turn...")
		if m.rollAccuracy(p) {
			hp := p.getHealth()
			m.dealDamage(p)
			fmt.Printf("%s dealt %d damage. \n", m.getName(), hp-p.getHealth())
			fmt.Printf("Player health remaining: %d \n", p.getHealth())
			if p.isDead() {
				fmt.Println("Battle finished. You died.")
				return
			}
		} else {
			fmt.Println("Attack missed.")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
