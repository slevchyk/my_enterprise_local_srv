package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Trailer struct {
	Id        uint64    `json:"id"`
	ExtId     string    `json:"ext_id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	MaxWeight float32   `json:"max_weight"`
	PhotoPath string    `json:"photo_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
