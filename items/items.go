package items

import "labirynth/game/entity"

type Item struct {
	entity.Entity
	impact int
}

type ItemLabyrinthProperties struct {
	Items []Item
}
