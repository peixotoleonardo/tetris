package internals

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240
	sampleRate   = 48000

	bytesPerSample = 4

	introLengthInSecond = 5
	loopLengthInSecond  = 120
)

type Game struct {
	sceneManager *SceneManager
	input        Input
	audio        Audio
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}

	g.input.Update()

	if err := g.audio.Update(); err != nil {
		return err
	}

	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}
