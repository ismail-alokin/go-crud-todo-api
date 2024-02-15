package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/config/logger"
)

func LogRequest(c *gin.Context) {
	logger.Println(c.Request.RequestURI)
	c.Next()
}
