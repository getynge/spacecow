// Package game describes the core constructs the game needs in order to perform more complex tasks.
// This package does not define any concrete types outside of Game, and keeps functions at a bare minimum,
// it primarily serves the purpose of creating the common components all the other packages will use to interact with
// one another.
package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func New() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
