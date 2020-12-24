package generic

import (
	"sort"

	"github.com/cbrand/ir-remote-backend/protocol"
)

// TheaterSorted joins a By function and a slice of Planets to be sorted.
type TheaterSorter struct {
	theaters []*protocol.Theater
	by       func(theater1, theater2 *protocol.Theater) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (sorter *TheaterSorter) Len() int {
	return len(sorter.theaters)
}

// Swap is part of sort.Interface.
func (sorter *TheaterSorter) Swap(i, j int) {
	sorter.theaters[i], sorter.theaters[j] = sorter.theaters[j], sorter.theaters[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (sorter *TheaterSorter) Less(i, j int) bool {
	return sorter.by(sorter.theaters[i], sorter.theaters[j])
}

// Sort runs the sort function of the provided theaters.
func (sorter *TheaterSorter) Sort(theaters []*protocol.Theater) {
	sorter.theaters = theaters
	sort.Sort(sorter)
}

// SortTheatersBy is used to sort theaters by a specified sort key.
func SortTheatersBy(sortedFunction func(theater1, theater2 *protocol.Theater) bool) *TheaterSorter {
	return &TheaterSorter{
		by: sortedFunction,
	}
}

// SortTheatersBySortKey sorts theaters by their internal sorting method
func SortTheatersBySortKey(theaters []*protocol.Theater) {
	sorter := SortTheatersBy(func(theater1, theater2 *protocol.Theater) bool {
		return theater1.SortKey < theater2.SortKey
	})
	sorter.Sort(theaters)
}
