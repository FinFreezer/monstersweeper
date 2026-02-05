package monstersweeper

import 
(
	"errors"
	"math/rand"
	"math"
	//"time"
	"fmt"
)

type Field struct {
	Tiles []Tile
	Grid map[int]map[int]*Tile
}

func (f *Field) calcTilePos() {
	tiles := []Tile{}
	var coord_x float32 = 0.0 + EDGE_MARGIN / 2
	var coord_y float32 = 0.0 + EDGE_MARGIN / 2
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++  {
			t := Tile{
					OriginX: coord_x,
					OriginY: coord_y,
					Width: TILE_SIZE_X,
					Height: TILE_SIZE_Y,
					IsMine: false,
					GridX: float32(i+1),
					GridY: float32(j+1),
				}
			if f.Grid[int(t.GridX)] == nil {
				f.Grid[int(t.GridX)] = make(map[int]*Tile)
			}
			f.Grid[int(t.GridX)][int(t.GridY)] = &t
			coord_x += (TILE_SIZE_X + TILE_MARGIN)
			tiles = append(tiles, t)
		}
		coord_y += (TILE_SIZE_Y + TILE_MARGIN)
		coord_x = 0.0 + EDGE_MARGIN / 2
	}
	f.Tiles = tiles
}

func InitField () (*Field, error) {
	f := Field{
		Tiles: []Tile{}, 
		Grid: make(map[int]map[int]*Tile),
	}
	f.calcTilePos()
	f.addMines()
	if len(f.Tiles) == 0 {
		return &f, errors.New("Field initialization failed.")
	}
	return &f, nil
}

func (f *Field) addMines() {
	//seedProd := time.Now().UnixNano()
	var seedDebug int64 = 1234
	r := rand.New(rand.NewSource(seedDebug))
	minesPos := make(map[int]map[int]bool)
	mines := []Mine{}

	for i := 0; i < int(math.Round(64 * 0.15)); i++ {
		randomX := r.Intn(8) + 1
		randomY := r.Intn(8) + 1

		if minesPos[randomX] == nil {
			minesPos[randomX] = make(map[int]bool)
		}

		for {
			if minesPos[randomX][randomY] {
				randomX = r.Intn(8) + 1
				randomY = r.Intn(8) + 1
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
	for _, mine := range(mines) {
		fmt.Println(mine)
		f.Grid[int(mine.posX)][int(mine.posY)].IsMine = true
	}
}

func (f *Field) FindClickedTile(coord_x, coord_y int) (tile *Tile) {
	for _, tile := range(f.Tiles) {
		if coord_x >= int(tile.OriginX) && coord_x <= int(tile.OriginX) + int(tile.Width) {
			if coord_y >= int(tile.OriginY) && coord_y <= int(tile.OriginY) + int(tile.Height) {
				fmt.Printf("Tile found at grid: %0.1f, %0.1f", tile.GridX, tile.GridY)
				return &tile
			}
		}
	}
	return nil
}