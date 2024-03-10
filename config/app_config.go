package config

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	vendor       string
	username     string
	password     string
	url          string
	databaseName string

	// Database source with username and password to which
	// open or connect call has to be done
	Source string
}

func (conf *DatabaseConfig) GenerateDatabaseSource() error {
	if conf.vendor == "" {
		return fmt.Errorf("database vendor name cannot be empty")
	}

	switch conf.vendor {
	case "postgres":
		connectionString := fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=disable",
			conf.username,
			conf.password,
			conf.url,
			conf.databaseName,
		)
		conf.Source = connectionString

	default:
		return fmt.Errorf("database vendor %s not supported", conf.vendor)
	}

	return nil
}

type AppConfig struct {
	Port             int
	AccessSecretKey  string
	RefreshSecretKey string
	DatabaseConf     DatabaseConfig
}

var DefaultConfig = AppConfig{}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*AppConfig, error) {
	config := AppConfig{}

	viper.AddConfigPath("/etc/secrets")
	viper.AddConfigPath(".")

	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	viper.BindEnv("app.port", "PORT")
	viper.AutomaticEnv()

	viper.SetDefault("app.port", 8080)

	if err := viper.ReadInConfig(); err != nil {
		if confErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// Ignoring the error as AutomaticEnv will load...
			log.Error("application.yml file not found", "message", confErr.Error())
			log.Info("don't worry, we'll try to populate config with global env variables")
		} else {
			return nil, err
		}
	}

	config.Port = viper.GetInt("app.port")
	config.AccessSecretKey = viper.GetString("app.access_secret_key")
	config.RefreshSecretKey = viper.GetString("app.refresh_secret_key")

	config.DatabaseConf.vendor = viper.GetString("database.vendor")
	config.DatabaseConf.username = viper.GetString("database.username")
	config.DatabaseConf.password = viper.GetString("database.password")
	config.DatabaseConf.url = viper.GetString("database.url")
	config.DatabaseConf.databaseName = viper.GetString("database.database_name")

	// TODO: validate that every field here is populated

	err := config.DatabaseConf.GenerateDatabaseSource()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
