package tasks

import (
	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/api/handlers"
	"github.com/ismail-alokin/go-crud-api-todo/config/db"
)

func CreateTask(c *gin.Context) {
	var requestBody struct {
		Task TaskModel `json:"task"`
	}

	err := c.BindJSON(&requestBody)
	handlers.CheckHttpBadRequest(c, err)

	params := map[string]interface{}{
		"title":       requestBody.Task.Title,
		"description": requestBody.Task.Description,
		"due_date":    requestBody.Task.Due_date,
	}

	var resultTask TaskModel

	edgeDbErr := db.DbClient.QuerySingle(c.Request.Context(), InsertTaskQuery, &resultTask, params)
	if handlers.CheckDbQueryFailure(c, edgeDbErr) {
		return
	}

	response := map[string]interface{}{
		"message": "Task added successfully",
		"task": map[string]interface{}{
			"id": resultTask.ID,
		},
	}
	handlers.SendSuccessJSONResponse(c, response)
}

func GetAllTasks(c *gin.Context) {
	var tasks []TaskModel

	edgeDbErr := db.DbClient.Query(c.Request.Context(), GetAllTasksQuery, &tasks)
	if handlers.CheckDbQueryFailure(c, edgeDbErr) {
		return
	}
	response := map[string]interface{}{
		"tasks": tasks,
	}

	handlers.SendSuccessJSONResponse(c, response)
}

func GetTask(c *gin.Context) {
	var task TaskModel

	taskId := c.Param("taskId")

	uuid, err := edgedb.ParseUUID(taskId)
	if handlers.CheckHttpBadRequest(c, err) {
		return
	}

	params := map[string]interface{}{
		"id": uuid,
	}

	edgeDbErr := db.DbClient.QuerySingle(c.Request.Context(), GetTaskQuery, &task, params)
	if handlers.CheckDbQueryFailure(c, edgeDbErr) {
		return
	}

	response := map[string]interface{}{
		"task": task,
	}
	handlers.SendSuccessJSONResponse(c, response)
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("taskId")

	uuid, uuidParseErr := edgedb.ParseUUID(taskId)
	if handlers.CheckHttpBadRequest(c, uuidParseErr) {
		return
	}

	var requestBody struct {
		Task TaskModel `json:"task"`
	}

	jsonParseErr := c.BindJSON(&requestBody)
	handlers.CheckHttpBadRequest(c, jsonParseErr)

	params := map[string]interface{}{
		"id":          uuid,
		"title":       requestBody.Task.Title,
		"description": requestBody.Task.Description,
		"due_date":    requestBody.Task.Due_date,
	}

	var resultTask TaskModel

	edgeDbErr := db.DbClient.QuerySingle(c.Request.Context(), UpdateTasksQuery, &resultTask, params)
	if handlers.CheckDbQueryFailure(c, edgeDbErr) {
		return
	}

	response := map[string]interface{}{
		"message": "Task edited successfully",
		"task": map[string]interface{}{
			"id": resultTask.ID,
		},
	}
	handlers.SendSuccessJSONResponse(c, response)
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")

	uuid, uuidParseErr := edgedb.ParseUUID(taskId)
	if handlers.CheckHttpBadRequest(c, uuidParseErr) {
		return
	}

	params := map[string]interface{}{
		"id": uuid,
	}
	var resultTask TaskModel

	edgeDbErr := db.DbClient.QuerySingle(c.Request.Context(), DeleteTasksQuery, &resultTask, params)
	if handlers.CheckDbQueryFailure(c, edgeDbErr) {
		return
	}

	response := map[string]interface{}{
		"message": "Task deleted successfully",
		"task": map[string]interface{}{
			"id": resultTask.ID,
		},
	}
	handlers.SendSuccessJSONResponse(c, response)

}
