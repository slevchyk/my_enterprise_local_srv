package models

type AppUserCniRecipientImport struct {
	Id          uint64 `json:"id"`
	ExtId       string `json:"ext_id"`	
	AppUserId   string `json:"app_user_id"`
	RecipientId string `json:"recipient_id"`
	IsActive    bool   `json:"is_active"`
}
