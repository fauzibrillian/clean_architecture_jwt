package config

type AppConfig struct {
	DBUSER string
	DBHOST string
	DBPASS string
	DBNAME string
	DBPORT uint
}

func InitConfig() *AppConfig {
	var response = new(AppConfig)

	return response
}
