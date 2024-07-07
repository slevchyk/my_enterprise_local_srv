package models

import (
	"time"
	
	"github.com/slevchyk/my_enterprise_local_srv/core"	
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen
type Vehicle struct {
	Id         uint64         `json:"id"`
	ExtId      string         `json:"ext_id"`
	Name       string         `json:"name"`
	IsDeleted  bool           `json:"is_deleted"`
	Length     float64        `json:"length"`
	Width      float64        `json:"width"`
	Height     float64        `json:"height"`
	MinWeight  float64        `json:"min_weight"`
	MaxWeight  float64        `json:"max_weight"`
	Comment    string         `json:"comment"`
	PhotoPath  string         `json:"photo_path"`
	NfcId      string         `json:"nfc_id"`
	DefTrailer *Trailer       `json:"def_trailer" objectbox:"link"`
	DefDriver  *ServiceWorker `json:"def_driver" objectbox:"link"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type VehicleImport struct {
	Id              uint64     `json:"id"`
	ExtId           string     `json:"ext_id"`
	Name            string     `json:"name"`
	IsDeleted       bool       `json:"is_deleted"`
	Length          core.Float `json:"length"`
	Width           core.Float `json:"width"`
	Height          core.Float `json:"height"`
	MinWeight       core.Float `json:"min_weight"`
	MaxWeight       core.Float `json:"max_weight"`
	Comment         string     `json:"comment"`
	PhotoPath       string     `json:"photo_path"`
	NfcId           string     `json:"nfc_id"`
	DefTrailerExtId string     `json:"def_trailer_ext_id"`
	DefDriverExtId  string     `json:"def_driver_ext_id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}