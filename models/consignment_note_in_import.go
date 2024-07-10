package models

type ConsignmentNoteInImport struct {
	Id             uint64                         `json:"srv_id"`
	ExtId          string                         `json:"ext_id"`
	AppId          string                         `json:"app_id"`
	Date           string                         `json:"date"`
	Number         string                         `json:"number"`
	OperationId    int                            `json:"operation_id"`
	StatusId       int                            `json:"status_id"`
	ExtNumber      string                         `json:"ext_number"`
	HarvestTypeId  string                         `json:"harvest_type_id"`
	VehicleId      string                         `json:"vehicle_id"`
	TrailerId      string                         `json:"trailer_id"`
	DepartureDate  string                         `json:"departure_date"`
	DriverId       string                         `json:"driver_id"`
	RecipientId    string                         `json:"recipient_id"`
	ManagerId      string                         `json:"manager_id"`
	SenderId       string                         `json:"sender_id"`
	AppUserId      string                         `json:"app_user_id"`
	Comment        string                         `json:"comment"`
	Gross          float64                        `json:"gross"`
	Tare           float64                        `json:"tare"`
	Net            float64                        `json:"net"`
	Humidity       float64                        `json:"humidity"`
	Weediness      float64                        `json:"weediness"`
	Oiliness       float64                        `json:"oiliness"`
	OilinessDry    float64                        `json:"oiliness_dry"`
	ErucicAcid     float64                        `json:"erucic_acid"`
	Glucosinolates float64                        `json:"glucosinolates"`
	Mycotoxins     float64                        `json:"mycotoxins"`
	Protein        float64                        `json:"protein"`
	ProteinDry     float64                        `json:"protein_dry"`
	Acid           float64                        `json:"acid"`
	Seals          string                         `json:"seals"`
	IsDeleted      bool                           `json:"is_deleted"`
	CreatedAt      string                         `json:"created_at"`
	UpdatedAt      string                         `json:"updated_at"`
	ChangedByApp   bool                           `json:"changed_by_app"`
	ChangedByAcc   bool                           `json:"changed_by_acc"`
	Goods          []GoodsConsignmentNoteInImport `json:"table_goods"`
}
