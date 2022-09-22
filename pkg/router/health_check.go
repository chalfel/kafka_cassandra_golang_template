package router

import "github.com/gin-gonic/gin"

func (a *App) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"up": true})
}
