package monstersweeper

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"io"
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
	MineText            = initFont(TILE_SIZE_Y * 0.8)
	FirstClick          = true
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
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(readFileToBytes()))
	if err != nil {
		log.Fatal(err)
	}
	return &TextType{
		faceSource,
		size,
	}
}

func readFileToBytes() []byte {
	file, err := os.Open("./resources/fonts/FantasyMagist.otf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size := info.Size()
	if size <= 0 {
		log.Fatal(err)
	}

	data := make([]byte, size)
	_, err = io.ReadFull(file, data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
