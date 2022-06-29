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
	MaxWeight core.Float `json:"max_weight"`
	PhotoPath string     `json:"photo_path"`
	NfcId     string     `json:"nfc_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
