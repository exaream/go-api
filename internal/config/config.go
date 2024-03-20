package config

import (
	"cmp"
	"os"

	"github.com/joho/godotenv"
)

const (
	defaultDBHost      = "127.0.0.1"
	defaultDBTz        = "UTC"
	defaultDBLang      = "en_US.UTF-8"
	defaultDBProtocol  = "tcp"
	defaultDBCollation = "utf8mb4_bin"
)

type Config struct {
	HTTPPort string
	DB       DB
	TestDB   DB
}

type DB struct {
	Host         string
	Port         string
	RootUser     string
	RootPassword string
	Name         string
	User         string
	Password     string

	// MySQL
	Timezone             string
	Language             string
	Protocol             string
	Collation            string
	AllowNativePasswords bool

	// PostgreSQL
	SSLMode string
}

func Load(dotenvPath string) error {
	return godotenv.Load(dotenvPath)
}

func Get() (*Config, error) {
	return &Config{
		HTTPPort: cmp.Or(os.Getenv("HTTP_PORT"), "8080"),
		DB: DB{
			Host:         cmp.Or(os.Getenv("DB_HOST"), defaultDBHost),
			Port:         cmp.Or(os.Getenv("DB_PORT"), "3306"),
			RootUser:     os.Getenv("DB_ROOT_USER"),
			RootPassword: os.Getenv("DB_ROOT_PASSWORD"),
			Name:         os.Getenv("DB_NAME"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),

			// MySQL
			Timezone:             cmp.Or(os.Getenv("DB_TIMEZONE"), defaultDBTz),
			Language:             cmp.Or(os.Getenv("DB_LANGUAGE"), defaultDBLang),
			Protocol:             cmp.Or(os.Getenv("DB_PROTOCOL"), defaultDBProtocol),
			Collation:            cmp.Or(os.Getenv("DB_COLLATION"), defaultDBCollation),
			AllowNativePasswords: os.Getenv("DB_ALLOW_NATIVE_PASSWORDS") == "true",

			// PostgreSQL
			SSLMode: os.Getenv("DB_SSL_MODE"),
		},
		TestDB: DB{
			Host:         cmp.Or(os.Getenv("TEST_DB_HOST"), defaultDBHost),
			Port:         cmp.Or(os.Getenv("TEST_DB_PORT"), "3307"),
			RootUser:     os.Getenv("TEST_DB_ROOT_USER"),
			RootPassword: os.Getenv("TEST_DB_ROOT_PASSWORD"),
			Name:         os.Getenv("TEST_DB_NAME"),
			User:         os.Getenv("TEST_DB_USER"),
			Password:     os.Getenv("TEST_DB_PASSWORD"),

			// MySQL
			Timezone:             cmp.Or(os.Getenv("TEST_DB_TIMEZONE"), defaultDBTz),
			Language:             cmp.Or(os.Getenv("TEST_DB_LANGUAGE"), defaultDBLang),
			Protocol:             cmp.Or(os.Getenv("TEST_DB_PROTOCOL"), defaultDBProtocol),
			Collation:            cmp.Or(os.Getenv("TEST_DB_COLLATION"), defaultDBCollation),
			AllowNativePasswords: os.Getenv("TEST_DB_ALLOW_NATIVE_PASSWORDS") == "true",

			// PostgreSQL
			SSLMode: os.Getenv("DB_SSL_MODE"),
		},
	}, nil
}
