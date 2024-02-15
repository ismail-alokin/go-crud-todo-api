package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ReturnTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now()})
}
