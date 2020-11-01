package entity

// Event represents all the entityEvent
type Event interface {
	getSource() Entity
}

type BaseEvent struct {
	source Entity
}

// AddEntityEvent adds entity to the game instance.
type AddEntityEvent struct {
	BaseEvent
	EntityToAdd Entity
}

type RemoveEntityEvent struct {
	BaseEvent
	EntityToRemove Entity
}

func (e *AddEntityEvent) getSource() Entity {
	return e.BaseEvent.source
}

func (e *RemoveEntityEvent) getSource() Entity {
	return e.BaseEvent.source
}