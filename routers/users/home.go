package users

import "github.com/gin-gonic/gin"

func InitUsersRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "users",
		})
	})
}
