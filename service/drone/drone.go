package drone

import (
	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
)

type DroneOpts struct {
}

func NewDroneService(opts DroneOpts) Service {
	return &DroneOpts{}
}

// GetDronePlane get drone total travel distance
func (d *DroneOpts) GetDronePlane(estate *estate.Estate, tree []*tree.Tree, maxDistance int) (int, Coordinate) {
	initTree := d.initTreeInPlotsAsMap(tree)
	drone := newDroneMovements(&estateOpts{width: estate.Width, length: estate.Length, treesLocation: initTree}, maxDistance)
	drone.startTrack()

	stats := drone.getStats()
	if maxDistance > 0 {
		return maxDistance, Coordinate{X: stats.x, Y: stats.y}
	}
	return (stats.totalHorizontalMovements + stats.totalVerticalMovements), Coordinate{X: stats.x, Y: stats.y}
}

// initTreeInPlotsAsMap convert trees array to map
func (d *DroneOpts) initTreeInPlotsAsMap(tree []*tree.Tree) map[coordinate]int {
	var treesLocation map[coordinate]int
	treesLocation = make(map[coordinate]int)

	for _, v := range tree {
		treesLocation[coordinate{x: v.XCoordinate, y: v.YCoordinate}] = v.Height
	}

	return treesLocation
}
