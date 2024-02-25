package users

import "github.com/gin-gonic/gin"

func InitUsersRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "users",
		})
	})

	routerGroup.POST("/register", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"data": "register",
		})
	})
}
