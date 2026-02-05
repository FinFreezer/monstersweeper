package main

import 
(
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"fmt"
	"image/color"
	d "github.com/finfreezer/monstersweeper/monstersweeper"
)

type Game struct{
	field 	         *d.Field
	input 	         *d.Input
	first_click      bool
	has_changed      bool
}

func (g *Game) Update() error {

	if g.field == nil {
		f, err := d.InitField()
		g.field = f
		if err != nil {
			fmt.Println("Error: %s", err)
		}
	}

	if g.input == nil {
		i := d.InitInput()
		g.input = i
	}

	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	squareClr := color.RGBA{0xA9, 0xAD, 0xD1, 0xff}

	if g.has_changed == true {
		if len(g.field.Tiles) != 0 {
			for _, t := range(g.field.Tiles) {
				vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, squareClr, false)
			}
		}
		g.has_changed = false
	}
	/*if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mouse pressed."), 585, 20)
	}*/
	g.input.Draw(screen)
	x, y := g.input.ReturnPos()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d, %d", x, y), 600, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return d.SCREENWIDTH, d.SCREENHEIGHT
}

func main() {
	ebiten.SetWindowSize(d.SCREENWIDTH, d.SCREENHEIGHT)
	ebiten.SetWindowTitle("Monstersweeper")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}