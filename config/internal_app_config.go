package config

type appConfig struct {
	port             int
	accessSecretKey  string
	refreshSecretKey string
	datasource       string
	runMigration     bool
}

func newAppConfig(port int, accessSecretKey string,
	refreshSecretKey string, datasource string, runMigration bool) (appConfig, error) {

	ac := appConfig{
		port:             port,
		accessSecretKey:  accessSecretKey,
		refreshSecretKey: refreshSecretKey,
		datasource:       datasource,
		runMigration:     runMigration,
	}

	return ac, nil
}

// A configuration structure which will be populated on running config.LoadConfig()
// This value must be populated only once while the program starts executing.
// The fields inside appConfig are private and only have getter methods, hence the
// fields once populated cannot be changed from any other part of the program.
var globalConfig appConfig
