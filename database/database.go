package database

import (
	"github.com/sbadung/togo/config"
	"github.com/sbadung/togo/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB(c config.DatabaseConfig) (*gorm.DB, error) {
	connectionString := config.PostgresConnectionString(c)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		return nil, err
	}

	return db, nil
}
