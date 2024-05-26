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
