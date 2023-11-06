package game

// Entity is any game object that has an ID
type Entity interface {
	// GetID is the ID of this entity, it must be unique and it must not change
	GetID() int
}
