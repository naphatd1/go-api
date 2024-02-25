package v2

import "github.com/gin-gonic/gin"

func InitHomeRoutesV2(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "2.0.0",
		})
	})
}
