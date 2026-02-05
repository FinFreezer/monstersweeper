package monstersweeper

func Init() {
	return
}

type Tile struct{
	OriginX float32
	OriginY float32
	Width float32
	Height float32
	GridX float32
	GridY float32
	IsMine bool
}

type Mine struct {
	isFlagged bool
	posX float32
	posY float32
}

func (m *Mine) returnPos() (x, y float32) {
	return m.posX, m.posY
}