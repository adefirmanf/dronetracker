package drone

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
	"github.com/stretchr/testify/assert"
)

func TestGetDronePlan(t *testing.T) {
	tests := []struct {
		name    string
		payload struct {
			estate      *estate.Estate
			trees       []*tree.Tree
			maxDistance int
		}
		wantDistance   int
		wantCoordinate Coordinate
	}{
		{
			name: "should return drone movements and return corrent movements",
			payload: struct {
				estate      *estate.Estate
				trees       []*tree.Tree
				maxDistance int
			}{
				estate: &estate.Estate{
					Length: 6,
					Width:  3,
				},
				trees: []*tree.Tree{
					{XCoordinate: 2, YCoordinate: 1, Height: 5},
					{XCoordinate: 3, YCoordinate: 1, Height: 3},
					{XCoordinate: 4, YCoordinate: 1, Height: 4},
				},
				maxDistance: 0,
			},
			wantDistance:   184,
			wantCoordinate: Coordinate{X: 6, Y: 3},
		},
		{
			name: "should return drone movements and return corrent movements with maximum distance",
			payload: struct {
				estate      *estate.Estate
				trees       []*tree.Tree
				maxDistance int
			}{
				estate: &estate.Estate{
					Length: 6,
					Width:  3,
				},
				trees: []*tree.Tree{
					{XCoordinate: 2, YCoordinate: 1, Height: 5},
					{XCoordinate: 3, YCoordinate: 1, Height: 3},
					{XCoordinate: 4, YCoordinate: 1, Height: 4},
				},
				maxDistance: 100,
			},
			wantDistance:   100,
			wantCoordinate: Coordinate{X: 3, Y: 2},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			gotDistance, gotCoordinate := NewDroneService(DroneOpts{}).GetDronePlane(v.payload.estate, v.payload.trees, v.payload.maxDistance)
			assert.Equal(t, gotDistance, v.wantDistance)
			assert.Equal(t, gotCoordinate, v.wantCoordinate)
		})
	}
}
