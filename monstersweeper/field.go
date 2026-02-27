package monstersweeper

import (
	"errors"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Field struct {
	Tiles           []*Tile
	Grid            map[int]map[int]*Tile
	MineTiles       []*Tile
	RevealedTiles   []*Tile
	Flags           int
	ActiveBattle    bool
	ActiveEncounter *Monster
	KeyTile         *Tile
}

func (f *Field) ReturnMineAmt() (total, left int) {
	revealed := 0
	for _, mine := range f.MineTiles {
		if mine.IsRevealed {
			revealed += 1
		}
	}
	return len(f.MineTiles), len(f.MineTiles) - revealed - f.Flags
}

func (f *Field) calcTilePos() {
	tiles := []*Tile{}
	var coord_x float32 = 0.0 + EDGE_MARGIN/2
	var coord_y float32 = 0.0 + EDGE_MARGIN/2
	for j := 0; j < FieldSize; j++ {
		for i := 0; i < FieldSize; i++ {
			t := &Tile{
				OriginX:       coord_x,
				OriginY:       coord_y,
				Width:         TILE_SIZE_X,
				Height:        TILE_SIZE_Y,
				IsMine:        false,
				GridX:         float32(i + 1),
				GridY:         float32(j + 1),
				IsRevealed:    false,
				AdjacentMines: 0,
			}
			if f.Grid[int(t.GridX)] == nil {
				f.Grid[int(t.GridX)] = make(map[int]*Tile)
			}
			f.Grid[int(t.GridX)][int(t.GridY)] = t
			coord_x += (TILE_SIZE_X + TILE_MARGIN)
			tiles = append(tiles, t)
		}
		coord_y += (TILE_SIZE_Y + TILE_MARGIN)
		coord_x = 0.0 + EDGE_MARGIN/2
	}
	f.Tiles = tiles
}

func InitField() (*Field, error) {
	f := Field{
		Tiles: []*Tile{},
		Grid:  make(map[int]map[int]*Tile),
	}
	f.calcTilePos()
	f.addMines()
	f.addMonsters()
	if len(f.Tiles) == 0 {
		return &f, errors.New("Field initialization failed.")
	}
	return &f, nil
}

func (f *Field) addMines() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	minesPos := make(map[int]map[int]bool)
	mines := []Mine{}
	ratio := 0.15 // Mine density
	if FieldSize >= 16 {
		ratio = 0.20
	}
	for i := 0; i < int(math.Round(float64(FieldSize)*float64(FieldSize)*ratio)); i++ {
		randomX := r.Intn(FieldSize) + 1
		randomY := r.Intn(FieldSize) + 1

		if minesPos[randomX] == nil {
			minesPos[randomX] = make(map[int]bool)
		}

		for {
			if minesPos[randomX][randomY] {
				randomX = r.Intn(FieldSize) + 1
				randomY = r.Intn(FieldSize) + 1
				continue

			} else {

				if minesPos[randomX] == nil {
					minesPos[randomX] = make(map[int]bool)
				}

				mine := Mine{isFlagged: false, posX: float32(randomX), posY: float32(randomY)}
				mines = append(mines, mine)
				minesPos[randomX][randomY] = true
				break
			}
		}
	}
	for _, mine := range mines {
		mineTile := f.Grid[int(mine.posX)][int(mine.posY)]
		mineTile.IsMine = true
		f.MineTiles = append(f.MineTiles, mineTile)
		f.includeMines(mineTile)
	}
	return
}

func (f *Field) FindClickedTile(coord_x, coord_y int, rightClick bool) (tile *Tile) {
	for _, tile := range f.Tiles {
		if coord_x >= int(tile.OriginX) && coord_x <= int(tile.OriginX)+int(tile.Width) {
			if coord_y >= int(tile.OriginY) && coord_y <= int(tile.OriginY)+int(tile.Height) {
				fmt.Printf("Tile found at grid: %0.1f, %0.1f\n", tile.GridX, tile.GridY)
				if !rightClick {
					f.tileClicked(tile)
					return tile

				} else {
					if !tile.IsFlagged {
						tile.IsFlagged = true
						f.Flags += 1
						return tile
					} else {
						tile.IsFlagged = false
						f.Flags -= 1
						return tile
					}
				}
			}
		}
	}
	return nil
}

func (f *Field) tileClicked(t *Tile) {
	if t.IsMine {
		if FirstClick {
			f.handleFirstClickMine(t)
			FirstClick = false
			fmt.Println("First click mine!")
			return
		}
		t.IsRevealed = true
		fmt.Println("Oops, you've hit a mine.")
		f.ActiveBattle = true
		f.ActiveEncounter = t.Encounter
		return
	}
	FirstClick = false
	f.revealTiles(t)
	return
}

func (f *Field) handleFirstClickMine(t *Tile) {
	t.IsMine = false
	directions := []struct{ stepX, stepY int }{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	coordX := int(t.GridX)
	coordY := int(t.GridY)
	for _, dir := range directions {
		next := f.Grid[coordX+dir.stepX][coordY+dir.stepY]

		if next != nil && !next.IsMine {
			next.IsMine = true

			for i, tile := range f.MineTiles { //Replace the old tile.
				if tile.IsMine == false {
					f.MineTiles[i] = next
					break
				}
			}
			break
		}
	}
	for _, tile := range f.Tiles {
		tile.AdjacentMines = 0
	}
	for _, mine := range f.MineTiles {
		f.includeMines(mine)
	}
	for _, mine := range f.MineTiles {
		mine.Encounter = nil
	}
	f.addMonsters()
	f.RevealedTiles = append(f.RevealedTiles, t)
	t.IsRevealed = true
	return
}

func (f *Field) revealTiles(t *Tile) {
	if t == nil {
		return
	}

	if t.IsRevealed {
		return
	}
	if t.IsMine {
		return
	}
	if t.AdjacentMines != 0 {
		t.IsRevealed = true
		f.RevealedTiles = append(f.RevealedTiles, t)
		return
	}
	f.RevealedTiles = append(f.RevealedTiles, t)
	t.IsRevealed = true

	directions := []struct{ stepX, stepY int }{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	coordX := int(t.GridX)
	coordY := int(t.GridY)

	for _, dir := range directions {
		next := f.Grid[coordX+dir.stepX][coordY+dir.stepY]
		f.revealTiles(next)
	}
	return
}

func (f *Field) includeMines(t *Tile) {
	if t == nil {
		return
	}
	if t.IsRevealed {
		return
	}

	coordX := int(t.GridX)
	coordY := int(t.GridY)
	//Clockwise starting from the previous tile on the X-axis
	directions := []struct{ stepX, stepY int }{
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
	}
	for _, dir := range directions {
		next := f.Grid[coordX+dir.stepX][coordY+dir.stepY]
		if next != nil && !next.IsMine {
			next.AdjacentMines += 1
		} else {
			continue
		}
	}
}

func (f *Field) addMonsters() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	encounters := len(f.MineTiles)

	keyHolder := r.Intn(encounters)

	for _, tile := range f.MineTiles {
		monster := r.Intn(5) + 1
		tile.Encounter = NewMonster(monster)
	}
	f.MineTiles[keyHolder].Encounter.KeyCarrier = true
	f.KeyTile = f.MineTiles[keyHolder]

	for _, tile := range f.MineTiles {
		if tile.Encounter == nil {
			errors.New("No encounter found.")
		}
	}
}

func (f *Field) locateKey(monstersKilled int) string {
	if float32(monstersKilled) > (float32(len(f.MineTiles)) * 0.75) {
		return fmt.Sprintf("The key appears to be at %d, %d\n", f.KeyTile.GridX, f.KeyTile.GridY)
	}

	if float32(monstersKilled) > (float32(len(f.MineTiles)) * 0.5) {
		return fmt.Sprintf("The key appears to be at row %d\n", f.KeyTile.GridX)
	}

	if float32(monstersKilled) > (float32(len(f.MineTiles)) * 0.25) {
		return fmt.Sprintf("The key appears to be at quadrant %d\n", f.getQuadrant())
	}

	if float32(monstersKilled) >= 1 {
		return "Your target wasn't carrying the key, keep looking."
	}

	return "Something went wrong."
}

func (f *Field) getQuadrant() int {
	if f.KeyTile.GridX < float32(FieldSize)/2 && f.KeyTile.GridY < float32(FieldSize)/2 {
		return 1
	}
	if f.KeyTile.GridX >= float32(FieldSize)/2 && f.KeyTile.GridY < float32(FieldSize)/2 {
		return 2
	}
	if f.KeyTile.GridX < float32(FieldSize)/2 && f.KeyTile.GridY >= float32(FieldSize)/2 {
		return 3
	}
	if f.KeyTile.GridX >= float32(FieldSize)/2 && f.KeyTile.GridY >= float32(FieldSize)/2 {
		return 4
	}
	errors.New("Unreachable quadrant.")
	return 0
}

func (f *Field) Draw(screen *ebiten.Image) {
	if len(f.Tiles) != 0 {
		for _, t := range f.Tiles {
			if t.IsRevealed && t.IsMine {
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
				continue
			}
			if t.IsRevealed {
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
				continue
			}
			if t.IsFlagged {
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
				continue
			}
			vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, t.Height, TileClrInit, false)
			//Add shading

			drawShaders(screen, t)
			drawCorners(screen, t)
		}
	}
	f.drawMineInfo(screen)
}

func (f *Field) drawField(screen *ebiten.Image) {

}

// Draws the ratio of flagged/revealed mines to total mines.
func (f *Field) drawMineInfo(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	face := &text.GoTextFace{
		Source: GeneralText.Source,
		Size:   GeneralText.Size,
	}
	edgeTile := f.Tiles[FieldSize-1]
	edge := edgeTile.OriginX + edgeTile.Width + EDGE_MARGIN
	op.GeoM.Translate(float64(edge), float64(EDGE_MARGIN))
	op.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	total, left := f.ReturnMineAmt()
	mineText := fmt.Sprintf("Mines left %d / %d", left, total)
	text.Draw(screen, mineText, face, op)
}

// Adds lighter and darker areas to give the tiles a sense of depth.
func drawShaders(screen *ebiten.Image, t *Tile) {
	vector.FillRect(screen, (t.OriginX + t.Width - ShaderSize), t.OriginY, ShaderSize, t.Height, TileClrInitDark, false)
	vector.FillRect(screen, t.OriginX, t.OriginY+t.Height-ShaderSize, t.Width-ShaderSize, ShaderSize, TileClrInitDark, false)
	vector.FillRect(screen, t.OriginX, t.OriginY, t.Width, ShaderSize, TileClrInitLight, false)
	vector.FillRect(screen, t.OriginX, t.OriginY+ShaderSize, ShaderSize, t.Height-2*ShaderSize, TileClrInitLight, false)
}

// Fills in the missing corners to make the shaders look nicer.
func drawCorners(screen *ebiten.Image, t *Tile) {
	drawTopCorner(screen, t.OriginX, t.OriginY)
	drawBottomCorner(screen, t.OriginX, t.OriginY)
}

func drawTopCorner(screen *ebiten.Image, beginX, beginY float32) {
	var path vector.Path
	drawOp := &vector.DrawPathOptions{}
	drawOp.ColorScale.ScaleWithColor(TileClrInitDark)
	pathOp := &vector.AddPathOptions{}
	pathOp.GeoM.Translate(float64(beginX+TILE_SIZE_X), float64(beginY))
	path.AddPath(&TopCornerPath, pathOp)
	vector.FillPath(screen, &path, nil, drawOp)
	return
}

func drawBottomCorner(screen *ebiten.Image, beginX, beginY float32) {
	var path vector.Path
	drawOp := &vector.DrawPathOptions{}
	drawOp.ColorScale.ScaleWithColor(TileClrInitLight)
	pathOp := &vector.AddPathOptions{}
	pathOp.GeoM.Translate(float64(beginX), float64(beginY+TILE_SIZE_Y-ShaderSize))
	path.AddPath(&BottomCornerPath, pathOp)
	vector.FillPath(screen, &path, nil, drawOp)
	return
}
