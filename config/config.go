package config

import "os"

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Ssl      string
	Timezone string
}

func PostgresConn(mode string) PostgresConfig {

	if mode == "test" {
		return PostgresConfig{
			Host:     "db",
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Dbname:   os.Getenv("DB_NAME_TEST"),
			Port:     "5432",
			Ssl:      "disable",
			Timezone: "Asia/Taipei",
		}
	}

	return PostgresConfig{
		Host:     "db",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME_DEV"),
		Port:     "5432",
		Ssl:      "disable",
		Timezone: "Asia/Taipei",
	}
}

func MigrationsPath() string {
	return os.Getenv("MIGRATION_PATH")
}
