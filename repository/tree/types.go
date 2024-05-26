package tree

import "time"

type Tree struct {
	ID          string
	EstateID    string
	XCoordinate int
	YCoordinate int
	Height      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TreeStats struct {
	Count  int
	Max    int
	Min    int
	Median float32
}
