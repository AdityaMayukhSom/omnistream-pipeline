package database

// PostgreSQLDatabase represents a PostgreSQL database
type PostgreSQLDatabase struct {
	ConnectionURL string
}

// GetCredentials retrieves credentials for PostgreSQL
func (db PostgreSQLDatabase) GetCredentials(username string) *Credentials {
	cred := Credentials{
		Username:  username,
		AuthToken: "abcd1234",
	}

	return &cred
}

func (db PostgreSQLDatabase) GetCoins(username string) *CoinDetails {
	coins := CoinDetails{
		Username: username,
		Coins:    10000,
	}

	return &coins
}
