package database

// SQLiteDatabase represents an SQLite database
type SQLiteDatabase struct {
	Filepath string
}

// GetCredentials retrieves credentials for SQLite
func (db SQLiteDatabase) GetCredentials(username string) *Credentials {
	cred := Credentials{
		Username:  username,
		AuthToken: "ijkl9012",
	}

	return &cred
}

func (db SQLiteDatabase) GetCoins(username string) *CoinDetails {
	coins := CoinDetails{
		Username: username,
		Coins:    10000,
	}

	return &coins
}
