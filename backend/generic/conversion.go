package generic

import "github.com/cbrand/ir-remote-backend/protocol"

// TheaterSliceToMap converts a slice of theaters to a map which keys are the respective theater ids.
func TheaterSliceToMap(theaters []*protocol.Theater) map[string]*protocol.Theater {
	theatersByID := map[string]*protocol.Theater{}
	for _, theater := range theaters {
		theatersByID[theater.GetId()] = theater
	}
	return theatersByID
}

// TheaterMapToSortedSlice turns a map into a slice which is sorted by the internal sort key before
// returning it.
func TheaterMapToSortedSlice(theaterMap map[string]*protocol.Theater) []*protocol.Theater {
	theaters := []*protocol.Theater{}
	for _, theater := range theaterMap {
		theaters = append(theaters, theater)
	}
	SortTheatersBySortKey(theaters)
	return theaters
}
