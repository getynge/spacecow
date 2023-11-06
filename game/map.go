package game

type Coordinate [2]float64

type Map interface {
	// GetDimensions one or both dimensions may be infinite, they must not be negative infinity or NaN
	GetDimensions() Coordinate
}
