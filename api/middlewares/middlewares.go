package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-crud-api-todo/config/logger"
)

func LogRequest(c *gin.Context) {
	logger.Printf("%v: %v", c.Request.Method, c.Request.RequestURI)
	c.Next()
}
