package models

import "time"

type Product struct {
	Title          string
	Slug           string
	Description    string
	Price          float32
	Images         []string
	ProductOptions []byte
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
