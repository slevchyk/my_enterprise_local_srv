package models

import (
	"errors"
	"time"

	"github.com/objectbox/objectbox-go/objectbox"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type AppUser struct {
	Id                  uint64    `json:"id"`
	ExtId               string    `json:"ext_id"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	Email               string    `json:"email"`
	Phone               string    `json:"phone"`
	Password            string    `json:"password"`
	PhotoPath           string    `json:"photo_path"`
	Token               string    `json:"token"`
	TokenExpirationDate time.Time `json:"-"`
	IsAdministrator     bool      `json:"is_administrator"`
	IsManualSelecting   bool      `json:"is_manual_selecting"`
	IsBlocked           bool      `json:"is_blocked"`
	IsFarm              bool      `json:"is_farm"`
	IsGasStation        bool      `json:"is_gas_station"`
	IsHarvesting        bool      `json:"is_harvesting"`
	IsPayDesk           bool      `json:"is_pay_desk"`
	IsWarehouse         bool      `json:"is_warehouse"`
	IsReports           bool      `json:"is_reports"`
	IsDictionaries      bool      `json:"is_dictionaries"`
	IsElevator          bool      `json:"is_elevator"`
	IsViewMode          bool      `json:"is_view_mode"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func GetAppUserByToken(ob *objectbox.ObjectBox, token string) (*AppUser, error) {

	box := BoxForAppUser(ob)
	query := box.Query(AppUser_.Token.Equals(token, true))
	aus, err := query.Find()
	query.Close()
	if err != nil {
		return nil, err
	}

	if len(aus) == 0 {
		return nil, errors.New("app user not found")
	} else if len(aus) > 1 {
		return nil, errors.New("app user token colision")
	}

	au := aus[0]

	if au.TokenExpirationDate.Before(time.Now().UTC()) {
		return nil, errors.New("token expire")
	}

	return aus[0], nil
}

func (au *AppUser) CopyToExport() *AppUser {
	return &AppUser{
		Id:        au.Id,
		ExtId:     au.ExtId,
		FirstName: au.FirstName,
		LastName:  au.LastName,
	}
}

type AppUserCniRecipient struct {
	Id        uint64    `json:"srv_id" objectbox:"id"`
	ExtId     string    `json:"ext_id" objectbox:"index, unique"`
	AppUser   *AppUser  `json:"app_user" objectbox:"link"`
	Recipient *Storage  `json:"recipient" objectbox:"link"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
