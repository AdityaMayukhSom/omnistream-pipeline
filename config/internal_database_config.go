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
	baseurl string, databaseName string) (databaseConfig, error) {
	if vendor == "" {
		return databaseConfig{}, fmt.Errorf("database vendor cannot be empty")
	}
	if username == "" {
		return databaseConfig{}, fmt.Errorf("database username cannot be empty")
	}
	// password can be empty in some cases, so commenting this out
	// if password == ""  {
	// 	return databaseConfig{}, fmt.Errorf("database vendor name cannot be empty")
	// }

	if baseurl == "" {
		return databaseConfig{}, fmt.Errorf("database url cannot be empty")
	}

	dc := databaseConfig{
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
