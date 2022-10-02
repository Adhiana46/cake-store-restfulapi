package entities

import (
	"encoding/json"
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

func (e Cake) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}
