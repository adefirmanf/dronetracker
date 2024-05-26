package drone

import (
	"log"

	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
)

type DroneOpts struct {
	Logging             bool
	HorizontalDirection string
}

func NewDroneService(opts DroneOpts) Service {
	return &DroneOpts{
		Logging:             opts.Logging,
		HorizontalDirection: opts.HorizontalDirection,
	}
}

type Service interface {
	GetDronePlane(estate *estate.Estate, tree []*tree.Tree) int
}

func (d *DroneOpts) GetDronePlane(estate *estate.Estate, tree []*tree.Tree) int {
	initTree := d.initTreeInPlotsAsMap(tree)
	drone := newDroneMovements(&estateOpts{width: estate.Width, length: estate.Length, treesLocation: initTree})
	drone.startTrack()

	stats := drone.getStats()
	log.Println(stats)
	return stats.totalHorizontalMovements + stats.totalVerticalMovements
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
