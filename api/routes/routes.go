package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/api/handlers"
	"github.com/ismail-alokin/go-crud-api-todo/api/middlewares"
	"github.com/ismail-alokin/go-crud-api-todo/pkg/tasks"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/time", handlers.ReturnTime)

	api := router.Group("/api")
	api.Use(middlewares.LogRequest)

	tasksApi := api.Group("/tasks")

	{
		tasksApi.GET("/", tasks.GetAllTasks)
		tasksApi.POST("/", tasks.CreateTask)

		tasksApi.GET("/:taskId", tasks.GetTask)
		tasksApi.PUT("/:taskId", tasks.UpdateTask)
		tasksApi.DELETE("/:taskId", tasks.DeleteTask)
	}

}
