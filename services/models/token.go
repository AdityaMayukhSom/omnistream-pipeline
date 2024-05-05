package models

import "time"

type Token struct {
	AccessToken   string
	RefreshToken  string
	AccessUuid    string
	RefreshUuid   string
	IssuedAt      time.Time
	AccessExpire  time.Time
	RefreshExpire time.Time
}
