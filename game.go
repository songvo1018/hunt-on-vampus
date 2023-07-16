package game

import labirinthMap "labirynth/game/labyrinthMap"

type Game struct {
	environment labirinthMap.MapEnvironment
}

func (g Game) getEnvironment() labirinthMap.MapEnvironment {
	return g.environment
}
