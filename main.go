package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sbadung/togo/config"
	"github.com/sbadung/togo/database"
	"github.com/sbadung/togo/todo"
)

func main() {
	r := gin.Default()

	serverConfig := config.LoadServerConfig()

	db, err := database.InitPostgresDB(config.LoadPostgresConfig())
	if err != nil {
		fmt.Println(todo.ErrDatabaseInitialization)
		return
	}

	serverAddress := fmt.Sprintf(":%d", serverConfig.ServerPort)
	r.Run(serverAddress)
}
