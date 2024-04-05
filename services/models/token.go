package models

type Token struct {
	AccessToken   string
	RefreshToken  string
	AccessUuid    string
	RefreshUuid   string
	AccessExpire  int64
	RefreshExpire int64
}
