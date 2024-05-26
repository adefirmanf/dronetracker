package types

import "errors"

var (
	ErrorEstateNotFound     = errors.New("estate not found")
	ErorrTreeOutOfBound     = errors.New("tree out of bound")
	ErrorTreeAlreadyPlanted = errors.New("tree already planted in same coordinate")
)
