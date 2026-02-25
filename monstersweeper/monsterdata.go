package monstersweeper

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func newImp(id int) *Monster {
	m := Monster{
		Name:          "Imp",
		MaxHealth:     5,
		Health:        5,
		Strength:      6,
		Dexterity:     10,
		Intelligence:  8,
		MonsterId:     id,
		AnimFrames:    make(map[string][]*ebiten.Image),
		CurrentFrame:  0,
		PrevFrameTime: time.Now(),
	}
	m.readSprites()
	return &m
}

func newSkeleton(id int) *Monster {
	m := Monster{
		Name:          "Skeleton",
		MaxHealth:     8,
		Health:        8,
		Strength:      10,
		Dexterity:     5,
		Intelligence:  5,
		MonsterId:     id,
		AnimFrames:    make(map[string][]*ebiten.Image),
		CurrentFrame:  0,
		PrevFrameTime: time.Now(),
	}
	m.readSprites()
	return &m
}

func newZombie(id int) *Monster {
	m := Monster{
		Name:          "Zombie",
		MaxHealth:     12,
		Health:        12,
		Strength:      8,
		Dexterity:     4,
		Intelligence:  2,
		MonsterId:     id,
		AnimFrames:    make(map[string][]*ebiten.Image),
		CurrentFrame:  0,
		PrevFrameTime: time.Now(),
	}
	m.readSprites()
	return &m
}

func newWitch(id int) *Monster {
	m := Monster{
		Name:          "Witch",
		MaxHealth:     4,
		Health:        4,
		Strength:      6,
		Dexterity:     10,
		Intelligence:  10,
		MonsterId:     id,
		AnimFrames:    make(map[string][]*ebiten.Image),
		CurrentFrame:  0,
		PrevFrameTime: time.Now(),
	}
	m.readSprites()
	return &m
}

func newOrc(id int) *Monster {
	m := Monster{
		Name:          "Orc",
		MaxHealth:     14,
		Health:        14,
		Strength:      10,
		Dexterity:     4,
		Intelligence:  2,
		MonsterId:     id,
		AnimFrames:    make(map[string][]*ebiten.Image),
		CurrentFrame:  0,
		PrevFrameTime: time.Now(),
	}
	m.readSprites()
	return &m
}

func (m *Monster) readSprites() {
	m.readAttackSprites()
	m.readHurtSprites()
	m.readDeathSprites()
	return
}

func (m *Monster) readAttackSprites() {
	sprites := []*ebiten.Image{
		SlimeSprites["slime-attack-0.png"],
		SlimeSprites["slime-attack-1.png"],
		SlimeSprites["slime-attack-2.png"],
		SlimeSprites["slime-attack-3.png"],
		SlimeSprites["slime-attack-4.png"],
	}
	m.AnimFrames["attack"] = sprites
	return
}

func (m *Monster) readHurtSprites() {
	sprites := []*ebiten.Image{
		SlimeSprites["slime-hurt-0.png"],
		SlimeSprites["slime-hurt-1.png"],
		SlimeSprites["slime-hurt-2.png"],
		SlimeSprites["slime-hurt-3.png"],
	}
	m.AnimFrames["hurt"] = sprites
	return
}

func (m *Monster) readDeathSprites() {
	sprites := []*ebiten.Image{
		SlimeSprites["slime-die-0.png"],
		SlimeSprites["slime-die-1.png"],
		SlimeSprites["slime-die-2.png"],
		SlimeSprites["slime-die-3.png"],
	}
	m.AnimFrames["death"] = sprites
	return
}
