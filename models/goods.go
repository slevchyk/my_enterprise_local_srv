package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Goods struct {
	Id        int64     `json:"id"`
	ExtId     string    `json:"ext_id"`
	Name      string    `json:"name"`
	Unit      *Unit     `json:"unit" objectbox:"link"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}