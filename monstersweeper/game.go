package monstersweeper

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	field        *Field
	input        *Input
	GameOver     bool
	StageClear   bool
	ContinueFlag bool
	lockInput    bool
	player       *Player
	ActiveBattle bool
}

func MsweeperGame() (*Game, error) {
	g := &Game{}
	if g.player == nil {
		p := InitPlayer()
		g.player = p
	}
	if g.input == nil {
		i := InitInput()
		g.input = i
	}
	if g.field == nil {
		f, err := InitField()
		g.field = f
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func (g *Game) Update() error {

	//Lock input when in autobattle.
	if !g.ActiveBattle {
		g.input.Update()
	}

	//Checks if a battle was started, the game-bound flag is used to avoid duplicates.
	if g.field != nil && g.field.ActiveBattle && !g.ActiveBattle && g.field.ActiveEncounter != nil {
		if g.field.ActiveEncounter.getHealth() < 0 {
			g.field.ActiveBattle = false
			g.field.ActiveEncounter = nil
			g.ActiveBattle = false
			return nil
		}
		g.ActiveBattle = true
		go func() {
			StartBattle(g.player, g.field.ActiveEncounter, g.field)
			g.field.ActiveBattle = false
			g.field.ActiveEncounter = nil
			g.ActiveBattle = false
		}()
	}

	if g.StageClear && g.player.HasKey() {
		if g.lockInput {
			//Avoid player accidentally clicking through
			//the end-of-stage screen instantly.
			time.Sleep(2 * time.Second)
			g.lockInput = false
		}

		//Check if the player clicked to continue.
		if !g.ContinueFlag {
			if g.input.IsActive() {
				g.ContinueFlag = true
			}
			return nil
		}
		g.newStage()
	}

	if g.GameOver {
		return nil
	}

	//Reads the cursor position on clicks.
	if g.input.IsActive() {
		posX, posY := g.input.ReturnPos()
		g.field.FindClickedTile(posX, posY, g.input.WasRightClick())
		g.input.ClearRightClick()
	}
	g.StageClear = g.checkStageClear()
	g.GameOver = g.checkGameOver()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.ActiveBattle {
		BattleDraw(screen)
		return
	}

	if g.field != nil {
		g.field.Draw(screen)
	}

	if g.input != nil {
		g.input.Draw(screen)
	}

	g.drawDebug(screen)

	if g.GameOver {
		g.drawGameOver(screen)
	}

	if g.StageClear && g.player.HasKey() {
		g.drawStageTransition(screen)
	}
	g.drawGameInfo(screen)

}

func (g *Game) drawGameInfo(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	f := &text.GoTextFace{
		Source: GeneralText.Source,
		Size:   GeneralText.Size,
	}
	edgeTile := g.field.Tiles[FieldSize-1]
	edge := edgeTile.OriginX + edgeTile.Width + EDGE_MARGIN
	op.GeoM.Translate(float64(edge), float64(EDGE_MARGIN)+40)
	op.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

	if g.player.HasKey() {
		text.Draw(screen, "Key found.", f, op)
	} else {
		text.Draw(screen, "No key found.", f, op)
	}
}

func (g *Game) drawStageTransition(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, SCREENWIDTH, SCREENHEIGHT, color.RGBA{50, 50, 50, 200}, false)
	op := &text.DrawOptions{}
	f := &text.GoTextFace{
		Source: GeneralText.Source,
		Size:   GeneralText.Size,
	}
	x, y := text.Measure("Stage Cleared, press left click to continue", f, 0)
	op.GeoM.Translate((SCREENHEIGHT - x), (SCREENHEIGHT-y)/2)
	op.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

	text.Draw(screen, "Stage Cleared", f, op)
	op.GeoM.Translate(-90, y)
	text.Draw(screen, "Press left click to continue", f, op)
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	x, y := g.input.ReturnPos()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d, %d", x, y), 600, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.0f", ebiten.ActualTPS()), 0, 20)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, SCREENWIDTH, SCREENHEIGHT, color.RGBA{50, 50, 50, 200}, false)
	op := &text.DrawOptions{}
	f := &text.GoTextFace{
		Source: MineText.Source,
		Size:   MineText.Size,
	}
	op.GeoM.Translate(float64(TILE_SIZE_Y*3), float64(TILE_SIZE_Y*4))
	op.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

	text.Draw(screen, "Game Over", f, op)
}

func (g *Game) checkGameOver() bool {
	if g.player.Health <= 0 {
		return true
	}
	return false
}

func (g *Game) checkStageClear() bool {
	if len(g.field.RevealedTiles) == (len(g.field.Tiles) - len(g.field.MineTiles)) {
		g.lockInput = true
		return true
	}
	return false
}

func (g *Game) newStage() {
	FirstClick = true
	g.ContinueFlag = false
	FieldSize += 2
	InitTile()
	g.player.Items["Key"] = 0
	g.player.MonstersKilled = 0
	g.player.Health = g.player.MaxHealth
	f, err := InitField()
	g.field = f
	if err != nil {
		log.Fatal(err)
	}
	return
}
