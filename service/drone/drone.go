package drone

import (
	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
)

type DroneOpts struct {
	Logging             bool
	HorizontalDirection string
}

func NewDrone(opts DroneOpts) DroneInterface {
	return &DroneOpts{
		Logging:             opts.Logging,
		HorizontalDirection: opts.HorizontalDirection,
	}
}

type DroneInterface interface {
	GetDronePlane(estate *estate.Estate, tree []tree.Tree) int
}

type plots struct {
	treeLocationCoordinate map[struct{ x, y int }]int
}

func (d *DroneOpts) GetDronePlane(estate *estate.Estate, tree []tree.Tree) int {
	initTree := d.initTreeInPlotsAsMap(tree)
	drone := newDroneMovements(&estateOpts{width: estate.Width, length: estate.Length, treesLocation: initTree})
	drone.startTrack()

	drone.getStats()
	return 0
}

func (d *DroneOpts) initTreeInPlotsAsMap(tree []tree.Tree) map[coordinate]int {
	var treesLocation map[coordinate]int
	treesLocation = make(map[coordinate]int)

	for _, v := range tree {
		treesLocation[coordinate{x: v.XCoordinate, y: v.YCoordinate}] = v.Height
	}

	return treesLocation
}
