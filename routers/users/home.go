package users

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/naphat/gob-api/controllers/user"
)

func InitUsersRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", usercontroller.GetAll)

	routerGroup.POST("/register", usercontroller.Register)
}
