package monstersweeper

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	DebugSeed           = 12345
	FieldSize           = 8
	Images              = loadImages()
	TileClrInit         = color.RGBA{0xA9, 0xAD, 0xD1, 0xff}
	TileClrInitDark     = color.RGBA{0x72, 0x78, 0xA8, 0xff}
	TileClrInitLight    = color.RGBA{0xBF, 0xC2, 0xDB, 0xff}
	TileClrRevealed     = color.RGBA{0xff, 0xff, 0xff, 0xff}
	TileClrMineRevealed = color.RGBA{0xC8, 0x28, 0x28, 0xff}
	MineImg             = Images["mine.png"]
	MineText            = initFont(float64(TILE_SIZE_Y) * 0.8)
	FirstClick          = true
	FlagImg             = Images["flag.png"]
	GeneralText         = initFont(32)
	TILE_SIZE_Y         = float32((GAME_AREA_HEIGHT - ((FieldSize - 1) * TILE_MARGIN)) / FieldSize)
	TILE_SIZE_X         = TILE_SIZE_Y
	TopCornerPath       = drawTopCorner()
	BottomCornerPath    = drawBottomCorner()
	ShaderSize          = TILE_SIZE_Y * 0.1
)

type TextType struct {
	Source *text.GoTextFaceSource
	Size   float64
}

type Actor interface {
	rollAccuracy(target Actor) bool
	dealDamage(target Actor)
	takeDamage(damage int)
	getDexterity() int
	isDead() bool
	getHealth() int
	getName() string
}

const (
	EDGE_MARGIN      = 20
	SCREENWIDTH      = 1280
	SCREENHEIGHT     = 960
	TILE_MARGIN      = 5
	GAME_AREA_WIDTH  = SCREENWIDTH - 2*EDGE_MARGIN
	GAME_AREA_HEIGHT = SCREENHEIGHT - 2*EDGE_MARGIN
)

func loadImages() map[string]*ebiten.Image {
	path := "resources/images/"
	info, err := os.Stat(path)
	images := make(map[string]*ebiten.Image)

	if !info.IsDir() {
		log.Fatal("Resources not found.")
	}
	if err != nil {
		log.Fatal(err)
	}
	contents, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range contents {
		contentType := "File"
		if content.IsDir() {
			continue
		}
		file, err := os.Open(filepath.Join(path, content.Name()))
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		images[content.Name()] = ebiten.NewImageFromImage(img)
		fmt.Printf("[%s] %s\n", contentType, content.Name())
	}
	return images
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
	file, err := os.Open("resources/fonts/FantasyMagist.otf")
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

func drawTopCorner() vector.Path {
	var path vector.Path
	drawOp := &vector.DrawPathOptions{}
	drawOp.ColorScale.ScaleWithColor(TileClrInitDark)
	path.MoveTo(0, 0)
	path.LineTo(0, ShaderSize)
	path.LineTo(-ShaderSize, ShaderSize)
	path.LineTo(0, 0)
	path.Close()
	return path
}

func drawBottomCorner() vector.Path {
	var path vector.Path
	drawOp := &vector.DrawPathOptions{}
	drawOp.ColorScale.ScaleWithColor(TileClrInitLight)
	path.MoveTo(0, 0)
	path.LineTo(0, ShaderSize)
	path.LineTo(ShaderSize, 0)
	path.LineTo(0, 0)
	path.Close()
	return path
}

func RollDice(faces, dice int) []int {

	/*pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Printf("RollDice called from %s at %s:%d with dice=%d\n",
			funcName, file, line, dice)
	}*/

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rolls := []int{}
	for i := 0; i < dice; i++ {
		result := r.Intn(faces) + 1
		rolls = append(rolls, result)
	}
	return rolls
}
