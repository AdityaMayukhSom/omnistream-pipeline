package database

func GetInstance(databaseType DatabaseType) DatabaseInterface {
	var databaseInstance DatabaseInterface

	switch databaseType {
	case POSTGRESQL:
		databaseInstance = PostgreSQLDatabase{
			ConnectionURL: "localhost:5432",
		}

	case SQLITE:
		databaseInstance = SQLiteDatabase{
			Filepath: "details.db",
		}

	default:
		databaseInstance = nil
	}

	return databaseInstance
}
