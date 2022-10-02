package entities

import (
	"time"
)

type Cake struct {
	ID          int
	Title       string
	Description string
	Rating      float32
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
