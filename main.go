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

	todoRepository := todo.NewPostgreSQLRepository(db)

	/* Aggiungere dati da peristere nel context */
	r.Use(func(c *gin.Context) {
		c.Set("repository", todoRepository)
		c.Next()
	})

	todo.SetupRoutes(r)

	serverAddress := fmt.Sprintf(":%d", serverConfig.ServerPort)
	r.Run(serverAddress)
}
