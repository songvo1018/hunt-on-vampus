package labirinthMap

import (
	"labirynth/game/items"
)

type MapEnvironment struct {
	labyrinthProperties     MapProperties
	userLabyrinthLocation   MapLocation
	npcLabyrinthLocation    MapLocation
	itemLabyrinthProperties items.ItemLabyrinthProperties
}

type MapLocation struct {
	Longitude int
	Latitude  int
}

type MapError string

func (e MapError) Error() string {
	return string(e)
}

var (
	ErrNewLocationAlreadyBusy = MapError("this place is busy")
	ErrActorTypeUndefined     = MapError("actor type undefined")
)

type MapProperties struct {
	mapLongitude map[int]bool
	mapLatitude  map[int]bool
}

func (l *MapEnvironment) getPosition(target string) (MapLocation, error) {
	switch target {
	case "user":
		return l.userLabyrinthLocation, nil
	case "item":
		//TODO: get itemPosition by Item Type?
	case "npc":
		return l.npcLabyrinthLocation, nil
	default:
		return MapLocation{}, ErrActorTypeUndefined
	}
	return MapLocation{}, ErrActorTypeUndefined
}

func (l *MapEnvironment) changePosition(target string, newLocation MapLocation) (MapLocation, error) {
	isLocationBusy := l.labyrinthProperties.mapLongitude[newLocation.Longitude] && l.labyrinthProperties.mapLatitude[newLocation.Latitude]
	switch target {
	case "user":
		if isLocationBusy {
			return MapLocation{}, ErrNewLocationAlreadyBusy
		}
		l.userLabyrinthLocation.Latitude = newLocation.Latitude
		l.userLabyrinthLocation.Longitude = newLocation.Longitude
		return l.userLabyrinthLocation, nil
	case "npc":
		if isLocationBusy {
			return MapLocation{}, ErrNewLocationAlreadyBusy
		}
		l.npcLabyrinthLocation.Latitude = newLocation.Latitude
		l.npcLabyrinthLocation.Longitude = newLocation.Longitude
		return l.npcLabyrinthLocation, nil
	case "item":
		//... ErrNewLocationAlreadyBusy
	default:
		return MapLocation{}, nil
	}
	return MapLocation{}, nil
}
