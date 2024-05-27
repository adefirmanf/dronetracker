package drone

import (
	"math"
)

type coordinate struct {
	x, y int
}
type estateOpts struct {
	// Define the estate size
	width  int
	length int

	// Trees location with height as value and coordinate as a key
	treesLocation map[coordinate]int
}

type droneMovements struct {
	estateOpts
	gapHeightLevel int
	groundLevel    int

	// Stats of the drone
	currentDirection string
	x                int
	y                int
	heightDrone      int

	totalVerticalMovements   int
	totalHorizontalMovements int
}

func newDroneMovements(estate *estateOpts) *droneMovements {
	return &droneMovements{
		estateOpts:     *estate,
		gapHeightLevel: 1,
		groundLevel:    0,

		currentDirection: "east",
		x:                1,
		y:                1,
		heightDrone:      0,

		totalVerticalMovements:   0,
		totalHorizontalMovements: 0,
	}
}

func (d *droneMovements) getStats() *droneMovements {
	return d
}

// Check the current plot, if the current plot is tree,
// Adjust the drone height based on tree level
// If the current plot doesn't have a tree, the drone should be set to the ground level
func (d *droneMovements) updateDroneWhenInCurrentPlot(currentTreeHeight int) {
	stats := d
	// But if the current plot have a difference 1 after the subtraction
	// Adjust the drone height based on tree level
	if (stats.heightDrone - currentTreeHeight) != 1 {
		droneLevelHeight := currentTreeHeight + d.gapHeightLevel
		// Difference between drone height and tree height to calculate the totalVertical distance
		travelDiff := math.Abs(float64(stats.heightDrone - (droneLevelHeight)))
		stats.totalVerticalMovements += int(travelDiff)
		// Set the drone level to the tree level
		stats.heightDrone = droneLevelHeight
	}
}

func (d *droneMovements) updateDroneToGroundLevel() {
	stats := d
	if stats.heightDrone > 0 {
		travelDiff := stats.heightDrone
		stats.totalVerticalMovements += travelDiff
		// Set the drone level to the ground level
		stats.heightDrone = d.groundLevel
	}
}

// This function is to check if the next plot is tree but the drone is at the lowest level than the next plot/tree
// In order to avoid the collision, we need to increase the drone height before moving to the next plot
func (d *droneMovements) updateDroneBeforeToNextPlot(nextTreeHeight int) {
	stats := d

	droneLevelHeight := nextTreeHeight + stats.gapHeightLevel
	if (droneLevelHeight) > stats.heightDrone {
		travelDiff := math.Abs(float64(stats.heightDrone - (droneLevelHeight)))
		stats.totalVerticalMovements += int(travelDiff)
		stats.heightDrone = droneLevelHeight
	}
}

func (d *droneMovements) startTrack() {
	esOpts := d.estateOpts
	stats := d
	for i := 1; i < esOpts.length*esOpts.width; i++ {
		// Start from east to west (Drone position |> )
		if stats.currentDirection == "east" {

			if currentTreeHeight, ok := stats.treesLocation[coordinate{x: stats.x, y: stats.y}]; ok {
				d.updateDroneWhenInCurrentPlot(currentTreeHeight)
			} else {
				d.updateDroneToGroundLevel()
			}
			// This condition is to check if the next plot is tree but the drone is at the lowest level than the next plot/tree
			// In order to avoid the collision, we need to increase the drone height before moving to the next plot
			if nextTreeHeight, ok := stats.treesLocation[coordinate{x: stats.x + 1, y: stats.y}]; ok {
				d.updateDroneBeforeToNextPlot(nextTreeHeight)
			}
			stats.totalHorizontalMovements += 10
			stats.x++
		}

		// Start from west to east (Drone position <| )
		if stats.currentDirection == "west" {
			if currentTreeHeight, ok := stats.treesLocation[coordinate{x: stats.x, y: stats.y}]; ok {
				d.updateDroneWhenInCurrentPlot(currentTreeHeight)
			} else {
				d.updateDroneToGroundLevel()
			}

			if nextTreeHeight, ok := stats.treesLocation[coordinate{x: stats.x - 1, y: stats.y}]; ok {
				d.updateDroneBeforeToNextPlot(nextTreeHeight)
			}
			stats.totalHorizontalMovements += 10
			stats.x--
		}

		// If drone already reached the east or west side, change the direction to north
		if stats.currentDirection == "east" && stats.x == esOpts.length || stats.currentDirection == "west" && stats.x == 1 {
			stats.currentDirection = "north"
			// Need to check in north position for the tree
			if nextTreeHeight, ok := stats.treesLocation[coordinate{x: stats.x, y: stats.y + 1}]; ok {
				d.updateDroneBeforeToNextPlot(nextTreeHeight)
			}
			continue
		}

		// When the drone already change the direction
		if stats.currentDirection == "north" {
			// Then need to move the drone to the next column (move: 1 step only)
			stats.y++
			stats.totalHorizontalMovements += 10

			// If the next column is already at the edge (east), change the direction to west
			if stats.x == esOpts.length {
				stats.currentDirection = "west"
			}

			// If the next column is not at the edge (west), change the direction to east
			if stats.x == 1 {
				stats.currentDirection = "east"
			}
			continue
		}
	}
}
