package monstersweeper

import (
	_ "embed"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	TileClrInit         = color.RGBA{0xA9, 0xAD, 0xD1, 0xff}
	TileClrRevealed     = color.RGBA{0xff, 0xff, 0xff, 0xff}
	TileClrMineRevealed = color.RGBA{0xC8, 0x28, 0x28, 0xff}
	MineImg             = loadImages()
	MineText            = initFont(24)
)

type TextType struct {
	Source *text.GoTextFaceSource
	Size   float64
}

const (
	EDGE_MARGIN      = 20
	SCREENWIDTH      = 1280
	SCREENHEIGHT     = 960
	TILE_MARGIN      = 5
	GAME_AREA_WIDTH  = SCREENWIDTH - 2*EDGE_MARGIN
	GAME_AREA_HEIGHT = SCREENHEIGHT - 2*EDGE_MARGIN
	//TILE_SIZE_X = (GAME_AREA_WIDTH - 6 * TILE_MARGIN / 2) / 8
	TILE_SIZE_X = TILE_SIZE_Y
	TILE_SIZE_Y = (GAME_AREA_HEIGHT - 6*TILE_MARGIN/2) / 8
)

func loadImages() *ebiten.Image {
	file, err := os.Open("./resources/images/mine.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	mine := ebiten.NewImageFromImage(img)
	return mine
}

func initFont(size float64) *TextType {
	return &TextType{}
}
