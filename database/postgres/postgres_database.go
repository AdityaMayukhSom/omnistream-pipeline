package postgresql

import "devstream.in/pixelated-pipeline/services/models"

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

func (psql *PostgresDatabase) FindUserByUsername(username string) {}
func (psql *PostgresDatabase) FindUserByEmail(email string)       {}

func (psql *PostgresDatabase) DeleteUserByUsername(username string) {}
func (psql *PostgresDatabase) DeleteUserByEmail(email string)       {}

func (psql *PostgresDatabase) CreateUser(user models.User) {}
func (psql *PostgresDatabase) UpdateUser(user models.User) {}

func (psql *PostgresDatabase) FindPostById(id string)              {}
func (psql *PostgresDatabase) FindPostsByUsername(username string) {}

func (psql *PostgresDatabase) DeletePostById(id string) {}

func (psql *PostgresDatabase) CreatePost(post models.Post) {}
func (psql *PostgresDatabase) UpdatePost(post models.Post) {}

func (psql *PostgresDatabase) CleanupDatabase() {
	// Db, err := Database.DB()
	// if err != nil {
	// 	// do something useful
	// }

	// Db.Close()

}
