package models

import "time"

type Session struct {
	RefreshToken string `json:"-"`
	ExpiresAt    int64  `json:"-"`
}

type AccessToken struct {
	Id        uint64    `json:"id"`
	ExpiresAt time.Time `json:"expires_at"`
}
