package labirinthMap

import (
	"fmt"
	"labirynth/game/entity"
	"labirynth/game/items"
	"labirynth/game/location"
)

type MapEnvironment struct {
	labyrinthProperties     MapProperties
	userLabyrinthLocation   location.MapLocation
	npcLabyrinthLocation    location.MapLocation
	itemLabyrinthProperties items.ItemLabyrinthProperties
}

type MapAroundEntities struct {
	Position location.MapLocation
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
	ErrActorTypeUndefined = MapError("actor type undefined")
)

type MapProperties struct {
	GridSize      int
	EntitiesOnMap map[location.MapLocation]entity.Entity
}

func (p *MapProperties) fillEmptyLocation() {

	// todo! configuration
	for x := 0; x < p.GridSize; x++ {
		for y := 0; y < p.GridSize; y++ {
			thisLocation := location.MapLocation{X: x, Y: y}
			entityOnLocation, ok := p.EntitiesOnMap[thisLocation]
			if !ok {
				p.EntitiesOnMap[thisLocation] = entity.Entity{Class: entity.PassType, IsBarrier: false, Location: thisLocation}
			} else if entityOnLocation.Location == thisLocation {
				entityOnLocation.Location = thisLocation
				p.EntitiesOnMap[thisLocation] = entityOnLocation
			}
		}
	}
}

func (m *MapEnvironment) GetPosition(target string) (location.MapLocation, error) {
	switch target {
	case "user":
		return m.userLabyrinthLocation, nil
	case "item":
		//TODO: get itemPosition by Item Type?
	case "npc":
		return m.npcLabyrinthLocation, nil
	default:
		return location.MapLocation{}, ErrActorTypeUndefined
	}
	return location.MapLocation{}, ErrActorTypeUndefined
}

func (m *MapAroundEntities) String() string {
	return fmt.Sprintf("north: %v, northEast: %v, east: %v, southEast: %v, south: %v, southWest: %v, west: %v, northWest: %v", m.nEntity, m.neEntity, m.eEntity, m.seEntity, m.sEntity, m.swEntity, m.wEntity, m.nwEntity)
}

func (m *MapEnvironment) ChangePosition(target string, newLocation location.MapLocation) (location.MapLocation, error) {
	isLocationBusy := m.labyrinthProperties.EntitiesOnMap[newLocation].Class != entity.PassType
	switch target {
	case "user":
		if isLocationBusy {
			return location.MapLocation{}, location.ErrNewLocationAlreadyBusy
		}
		m.userLabyrinthLocation = newLocation
		return m.userLabyrinthLocation, nil
	case "npc":
		if isLocationBusy {
			return location.MapLocation{}, location.ErrNewLocationAlreadyBusy
		}
		m.npcLabyrinthLocation = newLocation
		return m.npcLabyrinthLocation, nil
	case "item":
		//... ErrNewLocationAlreadyBusy
	default:
		return location.MapLocation{}, nil
	}
	return location.MapLocation{}, nil
}

func (m *MapEnvironment) lookAround(observerLocation location.MapLocation) MapAroundEntities {
	entities := MapAroundEntities{Position: observerLocation}
	// north
	north := location.MapLocation{X: observerLocation.X, Y: observerLocation.Y + 1}
	//fmt.Println(north)
	northEast := location.MapLocation{X: observerLocation.X + 1, Y: observerLocation.Y + 1}
	northWest := location.MapLocation{X: observerLocation.X - 1, Y: observerLocation.Y + 1}
	entities.nEntity, _ = m.getEntityOnLocation(north)
	entities.neEntity, _ = m.getEntityOnLocation(northEast)
	entities.nwEntity, _ = m.getEntityOnLocation(northWest)
	// south
	south := location.MapLocation{X: observerLocation.X, Y: observerLocation.Y - 1}
	southEast := location.MapLocation{X: observerLocation.X + 1, Y: observerLocation.Y - 1}
	southWest := location.MapLocation{X: observerLocation.X - 1, Y: observerLocation.Y - 1}
	entities.sEntity, _ = m.getEntityOnLocation(south)
	entities.seEntity, _ = m.getEntityOnLocation(southEast)
	entities.swEntity, _ = m.getEntityOnLocation(southWest)
	//west
	west := location.MapLocation{X: observerLocation.X - 1, Y: observerLocation.Y}
	entities.wEntity, _ = m.getEntityOnLocation(west)
	//east
	east := location.MapLocation{X: observerLocation.X + 1, Y: observerLocation.Y}
	entities.eEntity, _ = m.getEntityOnLocation(east)
	return entities
}

func (m *MapEnvironment) getEntityOnLocation(location location.MapLocation) (entity.Entity, error) {
	entityOnLocation, _ := m.labyrinthProperties.EntitiesOnMap[location]
	//todo get entities ids from store to check barrier or not

	if entityOnLocation.Class == entity.WallType {
		return entity.Entity{Class: entity.WallType, IsBarrier: true, Location: entityOnLocation.Location}, nil
	} else {
		return entity.Entity{Class: entity.PassType, IsBarrier: false, Location: entityOnLocation.Location}, nil
	}

}
