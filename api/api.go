package api

import (
	"time"
)

type Product struct {
	id          int
	name        string
	description string
	price       float64
	category    string
	created_at  time.Time
}
