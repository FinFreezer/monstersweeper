package monstersweeper

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func InitTile() { //re-calculate variables after field refreshes.
	TILE_SIZE_Y = float32((GAME_AREA_HEIGHT - ((FieldSize - 1) * TILE_MARGIN)) / FieldSize)
	TILE_SIZE_X = TILE_SIZE_Y
	MineText = initFont(float64(TILE_SIZE_Y) * 0.8)
	ShaderSize = TILE_SIZE_Y * 0.1
	TopCornerPath = calcTopCorner()
	BottomCornerPath = calcBottomCorner()
}

type Tile struct {
	OriginX       float32
	OriginY       float32
	Width         float32
	Height        float32
	GridX         float32
	GridY         float32
	IsMine        bool
	IsRevealed    bool
	AdjacentMines int
	IsFlagged     bool
	Encounter     *Monster
}

type Mine struct {
	isFlagged bool
	posX      float32
	posY      float32
}

func (m *Mine) returnPos() (x, y float32) {
	return m.posX, m.posY
}

func (t *Tile) DrawRevealedMine(screen *ebiten.Image) {
	vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, TileClrMineRevealed, false)
	op := &ebiten.DrawImageOptions{}
	rect := MineImg.Bounds()
	width := rect.Dx()
	heigth := rect.Dy()
	scaleX := float64((float64(TILE_SIZE_X) / float64(width))) * 0.8
	scaleY := float64((float64(TILE_SIZE_Y) / float64(heigth))) * 0.8
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(t.OriginX+(t.Width/10)), float64(t.OriginY+(t.Height/10)))
	screen.DrawImage(MineImg, op)
	return
}

func (t *Tile) DrawRevealedTile(screen *ebiten.Image) {
	vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, TileClrRevealed, false)
	if t.AdjacentMines > 0 {
		mineAmt := strconv.Itoa(t.AdjacentMines)
		op := &text.DrawOptions{}
		f := &text.GoTextFace{
			Source: MineText.Source,
			Size:   MineText.Size,
		}
		x, y := text.Measure(mineAmt, f, 0)
		op.GeoM.Translate(float64(t.OriginX+((t.Width-float32(x))/2)), float64(t.OriginY+(t.Height-float32(y))))
		op.ColorScale.ScaleWithColor(color.RGBA{0x00, 0x80, 0x00, 0xff})

		text.Draw(screen, mineAmt, f, op)
	}
	return
}

func (t *Tile) DrawFlag(screen *ebiten.Image) {
	vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, TileClrInit, false)
	drawShaders(screen, t)
	drawCorners(screen, t)
	op := &ebiten.DrawImageOptions{}
	rect := FlagImg.Bounds()
	width := rect.Dx()
	heigth := rect.Dy()
	scaleX := float64((float64(TILE_SIZE_X) / float64(width))) * 0.8
	scaleY := float64((float64(TILE_SIZE_Y) / float64(heigth))) * 0.8
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(t.OriginX+(t.Width/10)), float64(t.OriginY+(t.Height/10)))
	screen.DrawImage(FlagImg, op)
	return
}

func (t *Tile) DrawGenericTile(screen *ebiten.Image) {
	vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, TileClrInit, false)
	return
}
