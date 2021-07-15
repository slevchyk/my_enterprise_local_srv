package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type AppUser struct {
	Id           int64  `json:"id"`
	ExtId        string `json:"ext_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Token        string `json:"token"`
	IsBlocked    bool   `json:"is_blocked"`
	IsFarm       bool   `json:"is_farm"`
	IsGasStation bool   `json:"is_gas_station"`
	IsHarvesting bool   `json:"is_harvesting"`
	IsPayDesk    bool   `json:"is_pay_desk"`
	IsWarehouse  bool   `json:"is_warehouse"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
