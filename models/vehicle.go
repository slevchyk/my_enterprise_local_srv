package models

import (
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/core"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Vehicle struct {
	Id        uint64     `json:"id"`
	ExtId     string     `json:"ext_id"`
	Name      string     `json:"name"`
	IsDeleted bool       `json:"is_deleted"`
	Length    core.Float `json:"length"`
	Width     core.Float `json:"width"`
	Height    core.Float `json:"height"`
	MinWeight core.Float `json:"min_weight"`
	MaxWeight core.Float `json:"max_weight"`
	Comment   string     `json:"comment"`
	PhotoPath string     `json:"photo_path"`
	NfcId     string     `json:"nfc_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
