package monstersweeper

import
(
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"fmt"
)

type Input struct {
	mousePosX int
	mousePosY int
	mouseActive bool
}

func InitInput () (newInput *Input) {
	i := Input{}
	return &i
}

func (i *Input) ReturnPos() (x, y int) {
	return i.mousePosX, i.mousePosY
}

func (i *Input) IsActive() (bool) {
	if i.mouseActive {
		return true
	} else {
		return false
	}
}

func (i *Input) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		i.mousePosX, i.mousePosY = ebiten.CursorPosition()
		i.mouseActive = true
	} else {
		i.mouseActive = false
	}
}

func (i *Input) Draw(screen *ebiten.Image) {
	if (i.mouseActive) {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mouse pressed."), 585, 20)
	}
}