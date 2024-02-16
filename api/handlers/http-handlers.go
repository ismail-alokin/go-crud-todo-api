package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/config/logger"
)

func SendSuccessJSONResponse(c *gin.Context, jsonResposeData map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": jsonResposeData})
}

func CheckHttpBadRequest(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}

func HandleServerError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func CheckDbQueryFailure(c *gin.Context, err error) bool {
	if err != nil {
		logger.Printf("Error while executing edgeDB query", err)
	}
	return CheckHttpBadRequest(c, err)
}
