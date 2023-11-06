package ships

import "spacecow/game"

type Fighter struct {
}

func (f *Fighter) GetID() int {
	return 2
}

func (f *Fighter) Move(delta game.Coordinate) {
	return
}
