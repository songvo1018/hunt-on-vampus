package labirinthMap

import (
	"labirynth/game/entity"
	"labirynth/game/items"
)

type MapEnvironment struct {
	labyrinthProperties     MapProperties
	userLabyrinthLocation   MapLocation
	npcLabyrinthLocation    MapLocation
	itemLabyrinthProperties items.ItemLabyrinthProperties
}

//MapLocation это ширина = (Latitude) и долгода (Longitude) ||
type MapLocation struct {
	Longitude int
	Latitude  int
}

type MapAroundEntities struct {
	Position MapLocation
	nEntity  entity.Entity
	neEntity entity.Entity
	eEntity  entity.Entity
	seEntity entity.Entity
	sEntity  entity.Entity
	swEntity entity.Entity
	wEntity  entity.Entity
	nwEntity entity.Entity
}

type MapError string

func (e MapError) Error() string {
	return string(e)
}

var (
	ErrNewLocationAlreadyBusy = MapError("this place is busy")
	ErrActorTypeUndefined     = MapError("actor type undefined")
	ErrInvalidLocation        = MapError("invalid location for get entity")
)

type MapProperties struct {
	mapLongitude map[int]entity.Entity
	mapLatitude  map[int]entity.Entity
}

func (m *MapEnvironment) GetPosition(target string) (MapLocation, error) {
	switch target {
	case "user":
		return m.userLabyrinthLocation, nil
	case "item":
		//TODO: get itemPosition by Item Type?
	case "npc":
		return m.npcLabyrinthLocation, nil
	default:
		return MapLocation{}, ErrActorTypeUndefined
	}
	return MapLocation{}, ErrActorTypeUndefined
}

func (m *MapEnvironment) ChangePosition(target string, newLocation MapLocation) (MapLocation, error) {
	isLocationBusy := m.labyrinthProperties.mapLongitude[newLocation.Longitude] != entity.Pass && m.labyrinthProperties.mapLatitude[newLocation.Latitude] != entity.Pass
	switch target {
	case "user":
		if isLocationBusy {
			return MapLocation{}, ErrNewLocationAlreadyBusy
		}
		m.userLabyrinthLocation.Latitude = newLocation.Latitude
		m.userLabyrinthLocation.Longitude = newLocation.Longitude
		return m.userLabyrinthLocation, nil
	case "npc":
		if isLocationBusy {
			return MapLocation{}, ErrNewLocationAlreadyBusy
		}
		m.npcLabyrinthLocation.Latitude = newLocation.Latitude
		m.npcLabyrinthLocation.Longitude = newLocation.Longitude
		return m.npcLabyrinthLocation, nil
	case "item":
		//... ErrNewLocationAlreadyBusy
	default:
		return MapLocation{}, nil
	}
	return MapLocation{}, nil
}

func (m *MapEnvironment) lookAround(location MapLocation) MapAroundEntities {
	entities := MapAroundEntities{Position: location}
	// north
	entities.nEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude, location.Longitude - 1})
	entities.neEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude + 1, location.Longitude - 1})
	entities.nwEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude - 1, location.Longitude - 1})
	// south
	entities.sEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude, location.Longitude + 1})
	entities.seEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude + 1, location.Longitude + 1})
	entities.swEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude - 1, location.Longitude + 1})
	//west
	entities.wEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude - 1, location.Longitude})
	//east
	entities.eEntity, _ = m.getEntityOnLocation(MapLocation{location.Latitude + 1, location.Longitude})
	return entities
}

func (m *MapEnvironment) getEntityOnLocation(location MapLocation) (entity.Entity, error) {
	eLong := m.labyrinthProperties.mapLongitude[location.Longitude]
	eLat := m.labyrinthProperties.mapLatitude[location.Latitude]

	if eLat == entity.Wall && eLong == entity.Wall {
		//fmt.Printf("lat: long - %v:%v \n", location.Latitude, location.Longitude)
		//fmt.Printf("%v \n", entity.Wall)
		return entity.Wall, nil
	} else {
		//fmt.Printf("lat: long - %v:%v \n", location.Latitude, location.Longitude)
		//fmt.Printf("%v \n", entity.Pass)
		return entity.Pass, nil
	}
}
