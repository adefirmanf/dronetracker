package estate

import "time"

type Estate struct {
	ID        string
	Width     int
	Length    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
