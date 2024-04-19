package postgresql

type PostgresDatabase struct {
}

func (psql *PostgresDatabase) SetupDatabase() {
	// Db, err := sql.Open("pgx", config.DefaultConfig.DatabaseConf.Source)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database.")
	// }

	// gormDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: Db,
	// }), &gorm.Config{})

	// Database = gormDB
}

func (psql *PostgresDatabase) CleanupDatabase() {
	// Db, err := Database.DB()
	// if err != nil {
	// 	// do something useful
	// }

	// Db.Close()

}
