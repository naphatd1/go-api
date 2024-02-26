package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/naphat/gob-api/configs"
	users "github.com/naphat/gob-api/routers/users"
	v1 "github.com/naphat/gob-api/routers/v1"
	v2 "github.com/naphat/gob-api/routers/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT")) // listen and serve on 0.0.0.0:8080
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	configs.Connection()

	apiV1 := router.Group("/api/v1")
	apiV2 := router.Group("/api/v2")
	user := router.Group("/api")

	v1.InitHomeRoutesV1(apiV1)
	v2.InitHomeRoutesV2(apiV2)
	users.InitUsersRoutes(user)
	return router
}
