package database

type PostgresDatabase struct {
}

func NewPostgresDatabase() *PostgresDatabase {
	return &PostgresDatabase{}
}

func (pgdb *PostgresDatabase) SetupDatabase() {
	// Db, err := sql.Open("pgx", config.DefaultConfig.DatabaseConf.Source)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database.")
	// }

	// gormDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: Db,
	// }), &gorm.Config{})

	// Database = gormDB
}

func CleanupDatabase() {
	// Db, err := Database.DB()
	// if err != nil {
	// 	// do something useful
	// }

	// Db.Close()

}
