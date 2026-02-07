package main

import (
	"fmt"
	"log"

	d "github.com/finfreezer/monstersweeper/monstersweeper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	field       *d.Field
	input       *d.Input
	first_click bool
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
	if g.input.IsActive() {
		g.field.FindClickedTile(g.input.ReturnPos())
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if len(g.field.Tiles) != 0 {
		for _, t := range g.field.Tiles {
			if t.IsRevealed && t.IsMine {
				vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, d.TileClrMineRevealed, false)
				op := &ebiten.DrawImageOptions{}
				rect := d.MineImg.Bounds()
				width := rect.Dx()
				heigth := rect.Dy()
				scaleX := float64((float64(d.TILE_SIZE_X) / float64(width))) * 0.8
				scaleY := float64((float64(d.TILE_SIZE_Y) / float64(heigth))) * 0.8
				op.GeoM.Scale(scaleX, scaleY)
				op.GeoM.Translate(float64(t.OriginX+(t.Width/10)), float64(t.OriginY+(t.Height/10)))
				screen.DrawImage(d.MineImg, op)
				continue
			}
			if t.IsRevealed {
				vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, d.TileClrRevealed, false)
				continue
			}
			vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, d.TileClrInit, false)
		}
	}
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
