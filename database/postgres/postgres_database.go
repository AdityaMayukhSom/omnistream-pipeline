package postgresql

import (
	"database/sql"
	goDefaultLog "log"
	"os"
	"time"

	"devstream.in/pixelated-pipeline/config"
	"devstream.in/pixelated-pipeline/database/repositories"
	"devstream.in/pixelated-pipeline/services/models"

	"github.com/charmbracelet/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDatabase struct {
	gormDB *gorm.DB
}

func NewPostgresDatabase() *PostgresDatabase {
	sqlDB, err := sql.Open("pgx", config.GetDataSourceUri())
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.New(
			goDefaultLog.New(os.Stdout, "\r\n", goDefaultLog.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			},
		),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	database := &PostgresDatabase{
		gormDB: gormDB,
	}

	return database
}

func (psql *PostgresDatabase) Migrate() {
	psql.gormDB.AutoMigrate(
		&repositories.UserEntity{},
		&repositories.PostEntity{},
		&repositories.TokenEntity{},
		&repositories.CommentEntity{},
	)
}

func (psql *PostgresDatabase) FindUserByUsername(username string) {

}
func (psql *PostgresDatabase) FindUserByEmail(email string) {}

func (psql *PostgresDatabase) DeleteUserByUsername(username string) {}
func (psql *PostgresDatabase) DeleteUserByEmail(email string)       {}

func (psql *PostgresDatabase) CreateUser(user models.User) {}
func (psql *PostgresDatabase) UpdateUser(user models.User) {}

func (psql *PostgresDatabase) FindPostById(id string)              {}
func (psql *PostgresDatabase) FindPostsByUsername(username string) {}

func (psql *PostgresDatabase) DeletePostById(id string) {}

func (psql *PostgresDatabase) CreatePost(post models.Post) {}
func (psql *PostgresDatabase) UpdatePost(post models.Post) {}

func (psql *PostgresDatabase) Close() {
	Db, err := psql.gormDB.DB()
	if err != nil {
		// do something useful
		panic(err)
	}

	Db.Close()
}
