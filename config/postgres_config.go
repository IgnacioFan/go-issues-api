package config

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Ssl      string
	Timezone string
}

func NewPostgresConfig(mode string) *PostgresConfig {

	if mode == "test" {
		return &PostgresConfig{
			Host:     "db",
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Dbname:   os.Getenv("DB_NAME_TEST"),
			Port:     "5432",
			Ssl:      "disable",
			Timezone: "Asia/Taipei",
		}
	}

	return &PostgresConfig{
		Host:     "db",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME_DEV"),
		Port:     "5432",
		Ssl:      "disable",
		Timezone: "Asia/Taipei",
	}
}

func (this *PostgresConfig) Url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		this.User,
		this.Password,
		this.Host,
		this.Port,
		this.Dbname,
		this.Ssl,
	)
}

func (this *PostgresConfig) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		this.Host,
		this.User,
		this.Password,
		this.Dbname,
		this.Port,
		this.Ssl,
		this.Timezone,
	)
}
