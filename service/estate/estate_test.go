package estate

import (
	"log"
	"math"
	"testing"
)

func TestDroneTracker(t *testing.T) {
	type estate struct {
		width  int
		length int
	}
	type coordinate struct {
		x, y int
	}

	droneTotalDistance := 1

	plots := make(map[coordinate]int)
	plots[coordinate{x: 2, y: 1}] = 5
	plots[coordinate{x: 3, y: 1}] = 3
	plots[coordinate{x: 4, y: 1}] = 4
	plots[coordinate{x: 4, y: 2}] = 4
	plots[coordinate{x: 6, y: 2}] = 2
	plots[coordinate{x: 2, y: 2}] = 2
	// trace := make(map[string]coordinate)

	init := estate{length: 6, width: 3}

	// Start from east
	x := 1
	y := 1
	direction := "east"

	if direction == "west" {
		x = init.length
	}

	droneHeight := 0
	travelVerticalDistance := 0
	for i := 1; i < init.length*init.width; i++ {
		// Start from west to east
		if direction == "east" {

			// DRONE MOVEMENT
			// Check the current plot, if the current plot is tree.
			// Then the drone should be stay the level based on previous high
			if currentPlotHeight, ok := plots[coordinate{x: x, y: y}]; ok {
				// But if the current plot have a difference 1 after the subtraction
				// Adjust the drone height based on tree level
				if (droneHeight - currentPlotHeight) != 1 {
					travelDiff := math.Abs(float64(droneHeight - (currentPlotHeight + 1)))
					travelVerticalDistance += int(travelDiff)
					droneHeight = currentPlotHeight + 1
				}
				log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
			} else { // If the current plot is not tree, we can safely to reduce the drone height to the ground level (0)
				if droneHeight > 0 {
					groundLevel := 0
					travelDiff := math.Abs(float64(droneHeight - groundLevel))
					travelVerticalDistance += int(travelDiff)
					droneHeight = groundLevel
					log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
				}
			}
			// This condition is to check if the next plot is tree but the drone is at the lowest level than the next plot/tree
			// In order to avoid the collision, we need to increase the drone height before moving to the next plot
			if nextPlotHeight, ok := plots[coordinate{x: x + 1, y: y}]; ok {
				if (nextPlotHeight + 1) > droneHeight {
					travelDiff := math.Abs(float64(droneHeight - (nextPlotHeight + 1)))
					travelVerticalDistance += int(travelDiff)
					droneHeight = nextPlotHeight + 1
					log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
				}
			}
			x++

		}

		// Start from east to west
		if direction == "west" {
			// log.Printf("{x: %d, y: %d, height: %d, total_distance: %d}", x, y, 0, droneTotalDistance)

			// DRONE MOVEMENT
			// Check the current plot, if the current plot is tree.
			// Then the drone should be stay the level based on previous high
			if currentPlotHeight, ok := plots[coordinate{x: x, y: y}]; ok {
				// But if the current plot have a difference 1 after the subtraction
				// Adjust the drone height based on tree level
				if (droneHeight - currentPlotHeight) != 1 {
					travelDiff := math.Abs(float64(droneHeight - (currentPlotHeight + 1)))
					travelVerticalDistance += int(travelDiff)
					droneHeight = currentPlotHeight + 1
				}
				log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
			} else { // If the current plot is not tree, we can safely to reduce the drone height to the ground level (0)
				if droneHeight > 0 {
					groundLevel := 0
					travelDiff := math.Abs(float64(droneHeight - groundLevel))
					travelVerticalDistance += int(travelDiff)
					droneHeight = groundLevel
					log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
				}
			}
			// This condition is to check if the next plot is tree but the drone is at the lowest level than the next plot/tree
			// In order to avoid the collision, we need to increase the drone height before moving to the next plot
			if nextPlotHeight, ok := plots[coordinate{x: x - 1, y: y}]; ok {
				if (nextPlotHeight + 1) > droneHeight {
					travelDiff := math.Abs(float64(droneHeight - (nextPlotHeight + 1)))
					travelVerticalDistance += int(travelDiff)
					droneHeight = nextPlotHeight + 1
					log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
				}
			}
			x--
		}

		// If drone already reached the edge in the east, change direction to north
		if direction == "east" && x == init.length {
			direction = "north"
			if nextPlotHeight, ok := plots[coordinate{x: x, y: y + 1}]; ok {
				if (nextPlotHeight + 1) > droneHeight {
					travelDiff := math.Abs(float64(droneHeight - (nextPlotHeight + 1)))
					travelVerticalDistance += int(travelDiff)
					droneHeight = nextPlotHeight + 1
					log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
				}
			}
			// log.Printf("{x: %d, y: %d, height: %d, total_distance: %d}", x, y, 0, droneTotalDistance)
			continue
		}

		// If drone already reached the edge in the west, change direction to north
		if direction == "west" && x == 1 {
			direction = "north"
			if nextPlotHeight, ok := plots[coordinate{x: x, y: y + 1}]; ok {
				if (nextPlotHeight + 1) > droneHeight {
					travelDiff := math.Abs(float64(droneHeight - (nextPlotHeight + 1)))
					travelVerticalDistance += int(travelDiff)
					droneHeight = nextPlotHeight + 1
					log.Printf("|> (x: %d, y: %d) Drone height: %d travelVerticalDistace: %d", x, y, droneHeight, travelVerticalDistance)
				}
			}
			// log.Printf("{x: %d, y: %d, height: %d, total_distance: %d}", x, y, 0, droneTotalDistance)
			continue
		}

		// When the drone already change the direction
		if direction == "north" {
			// Then need to move the drone to the next column (move: 1 step only)
			y++
			droneTotalDistance += 10
			// After moving to the next column, the direction should be changed to west or east based on the condition

			// If the next column is already at the edge (east), change the direction to west
			if x == init.length {
				direction = "west"
			}

			// If the next column is not at the edge (west), change the direction to east
			if x == 1 {
				direction = "east"
			}
			continue
		}

	}
}
