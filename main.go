package main

import (
	"fmt"
	"image/color"
	"log"

	"strconv"

	d "github.com/finfreezer/monstersweeper/monstersweeper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	field    *d.Field
	input    *d.Input
	GameOver bool
}

func (g *Game) Update() error {
	if g.GameOver {
		return nil
	}
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
		posX, posY := g.input.ReturnPos()
		g.field.FindClickedTile(posX, posY, g.input.WasRightClick())
		g.input.ClearRightClick()
	}
	g.GameOver = g.checkGameOver()
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

				if t.AdjacentMines > 0 {
					mineAmt := strconv.Itoa(t.AdjacentMines)
					op := &text.DrawOptions{}
					f := &text.GoTextFace{
						Source: d.MineText.Source,
						Size:   d.MineText.Size,
					}
					x, y := text.Measure(mineAmt, f, 0)
					op.GeoM.Translate(float64(t.OriginX+((t.Width-float32(x))/2)), float64(t.OriginY+(t.Height-float32(y))))
					op.ColorScale.ScaleWithColor(color.RGBA{0x00, 0x80, 0x00, 0xff})

					text.Draw(screen, mineAmt, f, op)
				}
				continue
			}
			if t.IsFlagged {
				vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, d.TileClrInit, false)
				g.drawShaders(screen, t)
				g.drawCorners(screen, t)
				op := &ebiten.DrawImageOptions{}
				rect := d.FlagImg.Bounds()
				width := rect.Dx()
				heigth := rect.Dy()
				scaleX := float64((float64(d.TILE_SIZE_X) / float64(width))) * 0.8
				scaleY := float64((float64(d.TILE_SIZE_Y) / float64(heigth))) * 0.8
				op.GeoM.Scale(scaleX, scaleY)
				op.GeoM.Translate(float64(t.OriginX+(t.Width/10)), float64(t.OriginY+(t.Height/10)))
				screen.DrawImage(d.FlagImg, op)
				continue
			}
			vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, d.TileClrInit, false)
			//Add shading
			g.drawShaders(screen, t)
			g.drawCorners(screen, t)
		}

	}

	g.input.Draw(screen)
	x, y := g.input.ReturnPos()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d, %d", x, y), 600, 0)

	if g.GameOver {
		vector.FillRect(screen, 0, 0, d.SCREENWIDTH, d.SCREENHEIGHT, color.RGBA{50, 50, 50, 200}, false)
		op := &text.DrawOptions{}
		f := &text.GoTextFace{
			Source: d.MineText.Source,
			Size:   d.MineText.Size,
		}
		//x, y := text.Measure("Game Over", f, 0)
		op.GeoM.Translate(float64(d.TILE_SIZE_Y*3), float64(d.TILE_SIZE_Y*4))
		op.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

		text.Draw(screen, "Game Over", f, op)
	}

	op := &text.DrawOptions{}
	f := &text.GoTextFace{
		Source: d.GeneralText.Source,
		Size:   d.GeneralText.Size,
	}
	//x, y := text.Measure("Game Over", f, 0)
	edgeTile := g.field.Tiles[d.FieldSize-1]
	edge := edgeTile.OriginX + edgeTile.Width + d.EDGE_MARGIN
	op.GeoM.Translate(float64(edge), float64(d.EDGE_MARGIN))
	op.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	total, left := g.field.ReturnMineAmt()
	mineText := fmt.Sprintf("Mines left %d / %d", left, total)
	text.Draw(screen, mineText, f, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return d.SCREENWIDTH, d.SCREENHEIGHT
}

func (g *Game) checkGameOver() bool {
	if len(g.field.RevealedTiles) == (len(g.field.Tiles) - len(g.field.MineTiles)) {
		return true
	}
	return false
}
func (g *Game) drawShaders(screen *ebiten.Image, t *d.Tile) {
	vector.FillRect(screen, (t.OriginX + t.Width - 10), t.OriginY, 10, t.Height, d.TileClrInitDark, false)
	vector.FillRect(screen, t.OriginX, t.OriginY+t.Height-10, t.Width-10, 10, d.TileClrInitDark, false)
	vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, 10, d.TileClrInitLight, false)
	vector.FillRect(screen, t.OriginX, t.OriginY+10, 10, t.Height-20, d.TileClrInitLight, false)
}

func (g *Game) drawCorners(screen *ebiten.Image, t *d.Tile) {
	g.drawTopCorner(screen, t.OriginX+t.Width, t.OriginY)
	g.drawBottomCorner(screen, t.OriginX+t.Width, t.OriginY)
}

func (g *Game) drawTopCorner(screen *ebiten.Image, beginX, beginY float32) {
	drawOp := &vector.DrawPathOptions{}
	drawOp.ColorScale.ScaleWithColor(d.TileClrInitDark)
	var path vector.Path
	path.MoveTo(beginX, beginY)
	path.LineTo(beginX, beginY+10)
	path.LineTo(beginX-10, beginY+10)
	path.LineTo(beginX, beginY)
	path.Close()
	vector.FillPath(screen, &path, nil, drawOp)
	return
}

func (g *Game) drawBottomCorner(screen *ebiten.Image, beginX, beginY float32) {
	drawOp := &vector.DrawPathOptions{}
	var path vector.Path
	drawOp.ColorScale.ScaleWithColor(d.TileClrInitLight)
	beginX = beginX - d.TILE_SIZE_X
	path.MoveTo(beginX, beginY+d.TILE_SIZE_Y-10)
	path.LineTo(beginX+10, beginY+d.TILE_SIZE_Y-10)
	path.LineTo(beginX, beginY+d.TILE_SIZE_Y)
	path.LineTo(beginX, beginY+d.TILE_SIZE_Y-10)
	path.Close()
	vector.FillPath(screen, &path, nil, drawOp)
	return
}

func main() {
	ebiten.SetWindowSize(d.SCREENWIDTH, d.SCREENHEIGHT)
	ebiten.SetWindowTitle("Monstersweeper")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
