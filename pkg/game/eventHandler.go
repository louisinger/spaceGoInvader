package game

import (
	"github.com/louisinger/spaceGoInvader/pkg/entity"
	"errors"
)

// EventHandler applies the Event to the Game.
type EventHandler struct {
	game *Game
}

func newEventHandler(game *Game) EventHandler {
	return EventHandler{
		game: game,
	}
}

func (handler *EventHandler) handle(e entity.Event) error {
	if addIdentityEvent, ok := e.(*entity.AddEntityEvent); ok {
		return handler.handleAddEntityEvent(addIdentityEvent)
	}

	if removeIdentityEvent, ok := e.(*entity.RemoveEntityEvent); ok {
		return handler.handleRemoveEntityEvent(removeIdentityEvent)
	}

	return errors.New("Can't handle the event")
}

func (handler *EventHandler) handleAddEntityEvent(e *entity.AddEntityEvent) error {
	handler.game.Entities = append(handler.game.Entities, e.EntityToAdd)
	return nil
}

func (handler *EventHandler) handleRemoveEntityEvent(e *entity.RemoveEntityEvent) error {
	var indexToRemove int = -1
	for i, entity := range handler.game.Entities {
		if entity == e.EntityToRemove {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		return errors.New("Cannot find the index of the identity to remove")
	}

	indexOfLastElement := len(handler.game.Entities) - 1
	handler.game.Entities[indexToRemove] = handler.game.Entities[indexOfLastElement]
	handler.game.Entities = handler.game.Entities[:indexOfLastElement]

	return nil
}