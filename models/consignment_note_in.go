package models

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type ConsignmentNoteIn struct {
	Id                          uint64        `json:"id"`
	ExtId                       string       `json:"ext_id"`
	AppId                       string       `json:"app_id"`
	Date                        time.Time    `json:"date"`
	Number                      string       `json:"number"`
	HarvestType                 *HarvestType `json:"harvetst_type" objectbox:"link"`
	Vehicle                     *Vehicle     `json:"vehicle" objectbox:"link"`
	DepartureDate               time.Time    `json:"departure_date"`
	Driver                      *Person      `json:"driver" objectbox:"link"`
	Recipient                   *Storage     `json:"recipient" objectbox:"link"`
	Sender                      *Storage     `json:"sender" objectbox:"link"`
	MateriallyResponsiblePerson *AppUser     `json:"materially_responsible_person" objectbox:"link"`
	ResponsiblePerson           *AppUser     `json:"responsible_person" objectbox:"link"`
	Gross                       float64      `json:"gross"`
	Tare                        float64      `json:"tare"`
	Net                         float64      `json:"net"`
	Humidity                    float64      `json:"humidity"`
	Weediness                   float64      `json:"weediness"`
	IsDeleted                   bool         `json:"is_deleted"`
	CreatedAt                   time.Time    `json:"created_at"`
	UpdatedAt                   time.Time    `json:"updated_at"`
}
