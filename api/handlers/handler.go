package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/config/logger"
)

func ReturnTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now()})
}

func PanicHandler() {
	handlerError := recover()
	if handlerError != nil {
		logger.Printf("Cannot recover from panic", handlerError)
	}
}

func GinErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		logger.Printf("Error occured in router", err)
		HandleServerError(c, err)
		break
	}
}
