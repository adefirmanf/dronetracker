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
			estate *estate.Estate
			trees  []*tree.Tree
		}
		want int
	}{
		{
			name: "should return drone movements and return corrent movements",
			payload: struct {
				estate *estate.Estate
				trees  []*tree.Tree
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
			},
			want: 184,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := NewDroneService(DroneOpts{}).GetDronePlane(v.payload.estate, v.payload.trees)
			assert.Equal(t, got, v.want)
		})
	}
}
