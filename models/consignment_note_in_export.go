package models

type ConsignmentNoteInExport struct {
	Document   *ConsignmentNoteIn        `json:"document"`
	TableGoods []*GoodsConsignmentNoteIn `json:"table_goods"`
}
