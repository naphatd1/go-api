package users

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/naphat/gob-api/controllers/user"
	"github.com/naphat/gob-api/middlewares"
)

func InitUsersRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", middlewares.AuthJWT(), usercontroller.GetAll)
	routerGroup.POST("/register", middlewares.AuthJWT(), usercontroller.Register)
	routerGroup.POST("/login", usercontroller.Login)
	routerGroup.GET("/:id", middlewares.AuthJWT(), usercontroller.GetById)
	routerGroup.GET("/search", middlewares.AuthJWT(), usercontroller.SearchByFullname)
	//get Profile
	routerGroup.GET("/me", middlewares.AuthJWT(), usercontroller.GetProfile)
}
