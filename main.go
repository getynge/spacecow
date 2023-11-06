package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"spacecow/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("space cow")
	if err := ebiten.RunGame(game.New()); err != nil {
		panic(err) // if this happens we got some problems to solve
	}

}
