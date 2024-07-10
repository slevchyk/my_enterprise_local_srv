package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen
type Goods struct {
	Id               uint64    `json:"id"`
	ExtId            string    `json:"ext_id"`
	Name             string    `json:"name"`
	IsDeleted        bool      `json:"is_deleted"`
	Unit             *Unit     `json:"unit" objectbox:"link"`
	IsOiliness       bool      `json:"is_oiliness"`
	IsOilinessDry    bool      `json:"is_oiliness_dry"`
	IsErucicAcid     bool      `json:"is_erucic_acid"`
	IsGlucosinolates bool      `json:"is_glucosinolates"`
	IsMycotoxins     bool      `json:"is_mycotoxins"`
	IsProtein        bool      `json:"is_protein"`
	IsProteinDry     bool      `json:"is_protein_dry"`
	IsAcid           bool      `json:"is_acid"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
