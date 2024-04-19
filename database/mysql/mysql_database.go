package mysql

import "devstream.in/pixelated-pipeline/services/models"

type MysqlDatabase struct {
}

func (msql MysqlDatabase) SetupDatabase() {
	// Db, err := sql.Open("pgx", config.DefaultConfig.DatabaseConf.Source)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database.")
	// }

	// gormDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: Db,
	// }), &gorm.Config{})

	// Database = gormDB
}

func (msql MysqlDatabase) FindUserByUsername(username string) {}
func (msql MysqlDatabase) FindUserByEmail(email string)       {}

func (msql MysqlDatabase) DeleteUserByUsername(username string) {}
func (msql MysqlDatabase) DeleteUserByEmail(email string)       {}

func (msql MysqlDatabase) CreateUser(user models.User) {}
func (msql MysqlDatabase) UpdateUser(user models.User) {}

func (msql MysqlDatabase) FindPostById(id string)              {}
func (msql MysqlDatabase) FindPostsByUsername(username string) {}

func (msql MysqlDatabase) DeletePostById(id string) {}

func (msql MysqlDatabase) CreatePost(post models.Post) {}
func (msql MysqlDatabase) UpdatePost(post models.Post) {}

func (msql MysqlDatabase) CleanupDatabase() {
	// Db, err := Database.DB()
	// if err != nil {
	// 	// do something useful
	// }

	// Db.Close()
}
