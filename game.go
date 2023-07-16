package game

type Game struct {
	environment LabyrinthEnvironment
}

func (g Game) getEnvironment() LabyrinthEnvironment {
	return g.environment
}
