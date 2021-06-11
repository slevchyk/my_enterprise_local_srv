package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type GoodsConsignmentNoteIn struct {
	Id                int64              `json:"id"`
	ExtId             string             `json:"ext_id"`
	AppId             int64              `json:"app_id"`
	ConsignmentNoteIn *ConsignmentNoteIn `json:"-" objectbox:"link"`
	Subdivision       *Subdivision       `json:"subdivision" objectbox:"link"`
	GoodsGroup        *GoodsGroup        `json:"goods_group" objectbox:"link"`
	Goods             *Goods             `json:"goods" objectbox:"link"`
	Unit              *Unit              `json:"unit" objectbox:"link"`
	LoadingPercentage float32            `json:"loading_percentage"`
	Quantity          float32            `json:"quantity"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}