package entity

type Entity string

var (
	Wall      = Entity("wall")
	Pass      = Entity("pass")
	Door      = Entity("door")
	Collapsed = Entity("collapsed")
)
