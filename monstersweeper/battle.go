package monstersweeper

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	activePlayer   bool
	activeMonster  bool
	currentPlayer  *Player
	currentMonster *Monster
	field          *Field
	displayEnd     bool
	endTimer       time.Time
)

func StartBattle(p *Player, m *Monster, f *Field) string {
	currentPlayer = p
	currentMonster = m
	field = f
	displayEnd = false
	var debug string
	for {
		activePlayer = true
		currentPlayer.CurrentFrame = 0
		currentMonster.CurrentFrame = 0
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Player turn...")
		if p.rollAccuracy(m) {
			hp := m.getHealth()
			p.dealDamage(m)
			debug += fmt.Sprintf("Player dealt %d damage. \n", hp-m.getHealth())
			debug += fmt.Sprintf("%s health remaining: %d \n", m.getName(), m.getHealth())
			activePlayer = false
			if m.isDead() {
				playerWin()
				displayEnd = true
				endTimer = time.Now()
				waitEndDisplay()
				return debug
			}
		} else {
			fmt.Println("Attack missed.")
			activePlayer = false
		}
		activeMonster = true
		currentPlayer.CurrentFrame = 0
		currentMonster.CurrentFrame = 0
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Monster turn...")
		if m.rollAccuracy(p) {
			hp := p.getHealth()
			m.dealDamage(p)
			debug += fmt.Sprintf("%s dealt %d damage. \n", m.getName(), hp-p.getHealth())
			debug += fmt.Sprintf("Player health remaining: %d \n", p.getHealth())
			activeMonster = false
			if p.isDead() {
				monsterWin()
				debug += fmt.Sprintf("%s won with %d health left. \n", m.getName(), m.getHealth())
				debug += "\n\n"
				return debug
			}
		} else {
			fmt.Println("Attack missed.")
			activeMonster = false
		}
	}
}

func waitEndDisplay() {
	for {
		if displayEnd == false {
			return
		}
	}
}

func playerWin() {
	p := currentPlayer
	m := currentMonster
	p.MonstersKilled += 1
	if m.KeyCarrier {
		if p.Items == nil {
			p.Items = make(map[string]int)
		}
		p.Items["Key"] = 1
		fmt.Println("You found the key!")
	} else {
		fmt.Println(field.locateKey(p.MonstersKilled))
	}
	fmt.Println("Battle finished. Monster defeated.")
	healChance := RollDice(3, 1)
	if healChance[0] == 2 {
		fmt.Println("Looted a healing a potion! Healing to full.")
		p.Health = p.MaxHealth
	}
	return
}

func monsterWin() {
	fmt.Println("Battle finished. You died.")
	return
}

func BattleDraw(screen *ebiten.Image) {
	if displayEnd {
		drawEndScreen(screen)
	}
	if activePlayer {
		drawPlayerTurn(screen)
	}
	if activeMonster {
		drawMonsterTurn(screen)
	}
}

func drawEndScreen(screen *ebiten.Image) {
	keyFoundText := "You found the key! You can proceed to the next stage."
	op := &text.DrawOptions{}
	f := &text.GoTextFace{
		Source: GeneralText.Source,
		Size:   GeneralText.Size,
	}

	if !currentMonster.KeyCarrier {
		x, _ := text.Measure(field.locateKey(currentPlayer.MonstersKilled), f, 0)
		op.GeoM.Translate((SCREENWIDTH-x)/2, SCREENHEIGHT/2)
		text.Draw(screen, field.locateKey(currentPlayer.MonstersKilled), f, op)

	} else {
		x, _ := text.Measure(keyFoundText, f, 0)
		op.GeoM.Translate((SCREENWIDTH-x)/2, SCREENHEIGHT/2)
		text.Draw(screen, keyFoundText, f, op)
	}

	if time.Since(endTimer) > time.Duration(4*time.Second) {
		displayEnd = false
	}
}

func drawPlayerTurn(screen *ebiten.Image) {
	drawBattleText(screen)
	p := currentPlayer
	op := &ebiten.DrawImageOptions{}
	img := p.AnimFrames["attack"][p.CurrentFrame]
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(SCREENWIDTH/2-img.Bounds().Dx()), float64(SCREENHEIGHT/2-img.Bounds().Dy()))
	screen.DrawImage(img, op)
	if time.Since(p.PrevFrameTime) > time.Duration(200*time.Millisecond) {
		p.CurrentFrame += 1
		p.PrevFrameTime = time.Now()
		if p.CurrentFrame > len(p.AnimFrames["attack"])-1 {
			p.CurrentFrame = 0
		}
	}
	m := currentMonster
	op2 := &ebiten.DrawImageOptions{}
	img2 := m.AnimFrames["hurt"][m.CurrentFrame]
	op2.GeoM.Scale(2, 2)
	op2.GeoM.Translate(float64(SCREENWIDTH/2+img.Bounds().Dx()), float64(SCREENHEIGHT/2-img.Bounds().Dy()+25))
	screen.DrawImage(img2, op2)
	if time.Since(m.PrevFrameTime) > time.Duration(200*time.Millisecond) {
		m.CurrentFrame += 1
		m.PrevFrameTime = time.Now()
		if m.CurrentFrame > len(m.AnimFrames["hurt"])-1 {
			m.CurrentFrame = 0
		}
	}
	return
}

func drawMonsterTurn(screen *ebiten.Image) {
	drawBattleText(screen)
	p := currentPlayer
	op := &ebiten.DrawImageOptions{}
	img := p.AnimFrames["hurt"][p.CurrentFrame]
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(SCREENWIDTH/2-img.Bounds().Dx()), float64(SCREENHEIGHT/2-img.Bounds().Dy()))
	screen.DrawImage(img, op)
	if time.Since(p.PrevFrameTime) > time.Duration(200*time.Millisecond) {
		p.CurrentFrame += 1
		p.PrevFrameTime = time.Now()
		if p.CurrentFrame > len(p.AnimFrames["hurt"])-1 {
			p.CurrentFrame = 0
		}
	}
	m := currentMonster
	op2 := &ebiten.DrawImageOptions{}
	img2 := m.AnimFrames["attack"][m.CurrentFrame]
	op2.GeoM.Scale(2, 2)
	op2.GeoM.Translate(float64(SCREENWIDTH/2+img.Bounds().Dx()), float64(SCREENHEIGHT/2-img.Bounds().Dy()+25))
	screen.DrawImage(img2, op2)
	if time.Since(m.PrevFrameTime) > time.Duration(200*time.Millisecond) {
		m.CurrentFrame += 1
		m.PrevFrameTime = time.Now()
		if m.CurrentFrame > len(m.AnimFrames["attack"])-1 {
			m.CurrentFrame = 0
		}
	}
	return
}

func drawBattleText(screen *ebiten.Image) {

	op := &text.DrawOptions{}
	f := &text.GoTextFace{
		Source: GeneralText.Source,
		Size:   GeneralText.Size,
	}
	x, y := text.Measure("Battle in progess...", f, 0)
	op.GeoM.Translate((SCREENWIDTH-x)/2, 0)
	text.Draw(screen, "Battle in progess...", f, op)
	op.GeoM.Translate(0, y+5)
	playerStatus := fmt.Sprintf("Player health remaining: %d", currentPlayer.Health)
	text.Draw(screen, playerStatus, f, op)
	op.GeoM.Translate(0, y+5)
	monsterStatus := fmt.Sprintf("Monster health remaining: %d", currentMonster.Health)
	text.Draw(screen, monsterStatus, f, op)
}
