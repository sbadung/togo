package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServerConfig struct {
	ServerPort uint16 `json:"server_port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	Database string `json:"database"`
}

func PostgresConnectionString(c DatabaseConfig) string {
	// return fmt.Sprintf("host=%s user=%s password=%s dbname=%s", c.Host, c.User, c.Password, c.Database)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.Database)
	fmt.Println(connectionString)
	return fmt.Sprintf(connectionString)
}

func LoadPostgresConfig() DatabaseConfig {
	config := DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		Password: "postgres",
		User:     "postgres",
		Database: "postgres",
	}

	if postgresHost, exists := os.LookupEnv("POSTGRES_HOST"); exists {
		if postgresHost != "" {
			config.Password = postgresHost
		}
	}

	if postgresUser, exists := os.LookupEnv("POSTGRES_USER"); exists {
		if postgresUser != "" {
			config.User = postgresUser
		}
	}

	if postgresPassword, exists := os.LookupEnv("POSTGRES_PASSWORD"); exists {
		if postgresPassword != "" {
			config.Password = postgresPassword
		}
	}

	if postgresDatabase, exists := os.LookupEnv("POSTGRES_DB"); exists {
		if postgresDatabase != "" {
			config.Database = postgresDatabase
		}
	}

	return config
}

func LoadServerConfig() ServerConfig {
	config := ServerConfig{
		ServerPort: 8000,
	}

	/* Importare la configuratione da ENVIRONMENT VARIABLES */
	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			config.ServerPort = uint16(port)
		}
	}

	return config
}

func PostgresConnect(c DatabaseConfig) *gorm.DB {
	connectionString := PostgresConnectionString(c)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
