package tasks

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	fmt.Println("POST NEW TASKKKK")
}

func GetAllTasks(c *gin.Context) {
	fmt.Println("GET ALL TASKSSSSSSSSS")
}
