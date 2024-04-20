package config

import "fmt"

type databaseConfig struct {
	vendor       string
	username     string
	password     string
	baseurl      string
	databaseName string
}

func newDatabaseConfig(vendor string, username string, password string,
	baseurl string, databaseName string) (*databaseConfig, error) {
	if vendor == "" || username == "" || password == "" || baseurl == "" {
		return nil, fmt.Errorf("database vendor name cannot be empty")
	}

	dc := &databaseConfig{
		vendor:       vendor,
		username:     username,
		password:     password,
		baseurl:      baseurl,
		databaseName: databaseName,
	}

	return dc, nil

}

func (conf *databaseConfig) createDataSourceUri() (string, error) {
	var connectionString string = ""

	switch conf.vendor {
	case "postgres":
		connectionString = fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=disable",
			conf.username,
			conf.password,
			conf.baseurl,
			conf.databaseName,
		)

	default:
		return connectionString, fmt.Errorf("database vendor %s not supported", conf.vendor)
	}

	return connectionString, nil
}
