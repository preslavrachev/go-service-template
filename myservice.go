package myservice

import "time"

type Todo struct {
	Name          string
	IsCompleted   bool
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
