package models

type GoodsConsignmentNoteInImport struct {
	Id                uint64  `json:"srv_id"`
	ExtId             string  `json:"ext_id"`
	AppId             string  `json:"app_id"`
	SubdivisionId     string  `json:"subdivision_id"`
	GoodsGroupId      string  `json:"goods_group_Id"`
	GoodsId           string  `json:"goods_id"`
	UnitId            string  `json:"unit_id"`
	LoadingPercentage float32 `json:"loading_percentage"`
	Quantity          float32 `json:"quantity"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}
