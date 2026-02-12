package monstersweeper

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	mousePosX   int
	mousePosY   int
	mouseActive bool
	rightClick  bool
}

func InitInput() (newInput *Input) {
	i := Input{}
	return &i
}

func (i *Input) ReturnPos() (x, y int) {
	return i.mousePosX, i.mousePosY
}

func (i *Input) IsActive() bool {
	if i.mouseActive {
		return true
	} else {
		return false
	}
}

func (i *Input) WasRightClick() bool {
	return i.rightClick
}

func (i *Input) ClearRightClick() {
	i.rightClick = false
	return
}

func (i *Input) Update() {
	i.mouseActive = false

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		i.mousePosX, i.mousePosY = ebiten.CursorPosition()
		i.mouseActive = true
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		fmt.Println("Right Click.")
		i.mousePosX, i.mousePosY = ebiten.CursorPosition()
		i.mouseActive = true
		i.rightClick = true
	} else {
		i.mouseActive = false
	}
}

func (i *Input) Draw(screen *ebiten.Image) {
	if i.mouseActive {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mouse pressed."), 585, 20)
	}
}
