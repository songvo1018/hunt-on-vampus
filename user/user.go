package user

import (
	"labirynth/game/entity"
	"labirynth/game/items"
	labyrinthMap "labirynth/game/labyrinthMap"
)

type Properties struct {
	maxHp    int
	hp       int
	strength int
}

type User struct {
	entity.Entity
	Location       labyrinthMap.MapLocation
	Items          items.Item
	UserProperties Properties
}
