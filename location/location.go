package location

//MapLocation это ширина = (x) и долгода (y) ||
type MapLocation struct {
	X int
	Y int
}

type LocationError string

func (e LocationError) Error() string {
	return string(e)
}

var (
	NothereLocation           = MapLocation{0, 0}
	ErrNewLocationAlreadyBusy = LocationError("this place is busy")
	ErrInvalidLocation        = LocationError("invalid location for get entity")
)
