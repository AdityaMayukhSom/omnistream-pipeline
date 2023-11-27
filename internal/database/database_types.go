package database

type DatabaseType int

const (
	POSTGRESQL DatabaseType = iota
	MYSQL
	SQLITE
	REDIS
	MARIADB
	ORACLEDB
	MONGODB
	FIREBASE
)

type DatabaseInterface interface {
	GetCredentials(username string) *Credentials
	GetCoins(username string) *CoinDetails
}
