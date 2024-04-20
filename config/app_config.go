package config

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type appConfig struct {
	port             int
	accessSecretKey  string
	refreshSecretKey string
	datasource       string
}

func newAppConfig(port int, accessSecretKey string,
	refreshSecretKey string, datasource string) (appConfig, error) {

	ac := appConfig{
		port:             port,
		accessSecretKey:  accessSecretKey,
		refreshSecretKey: refreshSecretKey,
		datasource:       datasource,
	}

	return ac, nil
}

var globalConfig appConfig

// LoadConfig reads configuration from file or environment variables.
func LoadApplicationConfig() error {

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
			return err
		}
	}

	databaseConfig, err := newDatabaseConfig(
		viper.GetString("database.vendor"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.url"),
		viper.GetString("database.database_name"),
	)
	if err != nil {
		return err
	}

	datasource, err := databaseConfig.createDataSourceUri()
	if err != nil {
		return err
	}

	globalConfig, err = newAppConfig(
		viper.GetInt("app.port"),
		viper.GetString("app.access_secret_key"),
		viper.GetString("app.refresh_secret_key"),
		datasource,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetPort() int                { return globalConfig.port }
func GetAccessSecretKey() string  { return globalConfig.accessSecretKey }
func GetRefreshSecretKey() string { return globalConfig.refreshSecretKey }
func GetDataSourceUri() string    { return globalConfig.datasource }
