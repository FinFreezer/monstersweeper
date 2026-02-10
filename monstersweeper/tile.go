package monstersweeper

func InitTile() { //re-calculate variables after field refreshes.
	TILE_SIZE_Y = float32((GAME_AREA_HEIGHT - ((FieldSize - 1) * TILE_MARGIN)) / FieldSize)
	TILE_SIZE_X = TILE_SIZE_Y
	MineText = initFont(float64(TILE_SIZE_Y) * 0.8)
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
}

type Mine struct {
	isFlagged bool
	posX      float32
	posY      float32
}

func (m *Mine) returnPos() (x, y float32) {
	return m.posX, m.posY
}
