package main

import "time"

type TableA struct {
	ID        uint
	Column1   string
	Column2   int
	Column3   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
