package repositories

type TokenEntity struct {
	AccessToken   string `json:"accesstoken"`
	RefreshToken  string `json:"refreshtoken"`
	AccessUuid    string `json:"accessuuid"`
	RefreshUuid   string `json:"refreshuuid"`
	AccessExpire  int64  `json:"accessexpire"`
	RefreshExpire int64  `json:"refreshexpire"`
}

func (TokenEntity) TableName() string {
	return "tokens"
}
