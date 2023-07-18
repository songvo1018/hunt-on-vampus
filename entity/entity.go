package entity

import (
	"labirynth/game/location"
)

type Entity struct {
	Class     Type
	IsBarrier bool
	Location  location.MapLocation
}

type Type string

var (
	WallEntity      = Entity{WallType, true, location.NothereLocation}
	PassEntity      = Entity{PassType, false, location.NothereLocation}
	ItemEntity      = Entity{ItemType, false, location.NothereLocation}
	DoorEntity      = Entity{DoorType, true, location.NothereLocation}
	DestroyedEntity = Entity{DestroyedType, true, location.NothereLocation}

	//WallType is blocked way
	WallType = Type("wall")
	PassType = Type("pass")
	ItemType = Type("item")
	DoorType = Type("door")
	//DestroyedType may be something, and its will be blocked way
	DestroyedType = Type("collapsed")
)
