package repositories

import "time"

type TokenEntity struct {
	IssuedAt      time.Time `json:"issued_at"`
	AccessToken   string    `json:"access_token"`
	RefreshToken  string    `json:"refresh_token"`
	AccessUuid    string    `json:"access_uuid"`
	RefreshUuid   string    `json:"refresh_uuid"`
	AccessExpire  time.Time `json:"access_expire"`
	RefreshExpire time.Time `json:"refresh_expire"`
}

func (TokenEntity) TableName() string {
	return "tokens"
}
