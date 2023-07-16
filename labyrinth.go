package game

type LabyrinthEnvironment struct {
	labyrinthProperties     LabyrinthProperties
	userLabyrinthLocation   LabyrinthMapLocation
	npcLabyrinthLocation    LabyrinthMapLocation
	itemLabyrinthProperties ItemLabyrinthProperties
}

type LabyrinthMapLocation struct {
	Longitude int
	Latitude  int
}

type LabyrinthError string

func (e LabyrinthError) Error() string {
	return string(e)
}

var (
	ErrNewLocationAlreadyBusy = LabyrinthError("this place is busy")
	ErrActorTypeUndefined     = LabyrinthError("actor type undefined")
)

type LabyrinthProperties struct {
	mapLongitude map[int]bool
	mapLatitude  map[int]bool
}

func (l *LabyrinthEnvironment) getPosition(target string) (LabyrinthMapLocation, error) {
	switch target {
	case "user":
		return l.userLabyrinthLocation, nil
	case "item":
		//TODO: get itemPosition by Item Type?
	case "npc":
		return l.npcLabyrinthLocation, nil
	default:
		return LabyrinthMapLocation{}, ErrActorTypeUndefined
	}
	return LabyrinthMapLocation{}, ErrActorTypeUndefined
}

func (l *LabyrinthEnvironment) changePosition(target string, newLocation LabyrinthMapLocation) (LabyrinthMapLocation, error) {
	isLocationBusy := l.labyrinthProperties.mapLongitude[newLocation.Longitude] && l.labyrinthProperties.mapLatitude[newLocation.Latitude]
	switch target {
	case "user":
		if isLocationBusy {
			return LabyrinthMapLocation{}, ErrNewLocationAlreadyBusy
		}
		l.userLabyrinthLocation.Latitude = newLocation.Latitude
		l.userLabyrinthLocation.Longitude = newLocation.Longitude
		return l.userLabyrinthLocation, nil
	case "npc":
		if isLocationBusy {
			return LabyrinthMapLocation{}, ErrNewLocationAlreadyBusy
		}
		l.npcLabyrinthLocation.Latitude = newLocation.Latitude
		l.npcLabyrinthLocation.Longitude = newLocation.Longitude
		return l.npcLabyrinthLocation, nil
	case "item":
		//... ErrNewLocationAlreadyBusy
	default:
		return LabyrinthMapLocation{}, nil
	}
	return LabyrinthMapLocation{}, nil
}
