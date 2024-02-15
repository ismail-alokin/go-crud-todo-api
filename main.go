package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/api/routes"
	"github.com/ismail-alokin/go-crud-api-todo/config/db"
	"github.com/ismail-alokin/go-crud-api-todo/config/logger"
)

func main() {
	logger.Println("Starting server")
	defer db.DbClient.Close()

	router := gin.Default()
	routes.InitializeRoutes(router)

	router.Run(":8081")
}
