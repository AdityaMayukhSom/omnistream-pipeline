package models

type Token struct {
	AccessToken   string `json:"accesstoken"`
	RefreshToken  string `json:"refreshtoken"`
	AccessUuid    string `json:"accessuuid"`
	RefreshUuid   string `json:"refreshuuid"`
	AccessExpire  int64  `json:"accessexpire"`
	RefreshExpire int64  `json:"refreshexpire"`
}
