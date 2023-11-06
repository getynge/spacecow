package ships

import "spacecow/game"

type Ship interface {
	game.Entity
	Move(delta game.Coordinate)
}
