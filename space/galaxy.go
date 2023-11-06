package space

import "spacecow/game"

type Object struct {
	Instance game.Entity
	Location game.Coordinate
}

type Galaxy struct {
	dimensions game.Coordinate
}

func (g *Galaxy) GetDimensions() game.Coordinate {
	return g.dimensions
}
