package drone

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
)

func TestDrone(t *testing.T) {
	drone := NewDrone(DroneOpts{})
	drone.GetDronePlane(&estate.Estate{
		Length: 6,
		Width:  3,
	}, []tree.Tree{
		{XCoordinate: 2, YCoordinate: 1, Height: 5},
		{XCoordinate: 3, YCoordinate: 1, Height: 3},
		{XCoordinate: 4, YCoordinate: 1, Height: 4},
	})

}
