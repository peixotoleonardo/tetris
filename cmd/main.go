package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/peixotoleonardo/tetris/internals"
)

func main() {
	ebiten.SetWindowSize(internals.ScreenWidth * 2, internals.ScreenHeight * 2)

	ebiten.SetWindowTitle("Tetris")

	if err := ebiten.RunGame(&internals.Game{}); err != nil {
		log.Fatal(err)
	}
}
