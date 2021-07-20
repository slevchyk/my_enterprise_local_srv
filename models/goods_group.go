package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type GoodsGroup struct {
	Id        uint64  `json:"id"`
	ExtId     string `json:"ext_id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
