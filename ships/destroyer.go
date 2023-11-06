package ships

import "spacecow/game"

type Destroyer struct {
}

func (d *Destroyer) GetID() int {
	return 1
}

func (d *Destroyer) Move(delta game.Coordinate) {
	return
}
