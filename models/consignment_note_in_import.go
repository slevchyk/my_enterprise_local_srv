package models

type ConsignmentNoteInImport struct {
	Id              uint64                         `json:"srv_id"`
	ExtId           string                         `json:"ext_id"`
	AppId           string                         `json:"app_id"`
	Date            string                         `json:"date"`
	Number          string                         `json:"number"`
	HarvestTypeId   string                         `json:"harvest_type_id"`
	VehicleId       string                         `json:"vehicle_id"`
	DepartureDateId string                         `json:"departure_date"`
	DriverId        string                         `json:"driver_id"`
	RecipientId     string                         `json:"recipient_id"`
	SenderId        string                         `json:"sender_id"`
	AppUserId       string                         `json:"app_user_id"`
	Gross           float64                        `json:"gross"`
	Tare            float64                        `json:"tare"`
	Net             float64                        `json:"net"`
	Humidity        float64                        `json:"humidity"`
	Weediness       float64                        `json:"weediness"`
	IsDeleted       bool                           `json:"is_deleted"`
	CreatedAt       string                         `json:"created_at"`
	UpdatedAt       string                         `json:"updated_at"`
	ChangedByApp    bool                           `json:"changed_by_app"`
	ChangedByAcc    bool                           `json:"changed_by_acc"`
	Goods           []GoodsConsignmentNoteInImport `goods:"goods"`
}