package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naphat/gob-api/configs"
	"github.com/naphat/gob-api/models"
)

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "users",
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	result := configs.DB.Debug().Create(&user)

	//เช็คอีเมล์ซ้ำ
	// userExist :=

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error})
		return
	}

	c.JSON(201, gin.H{
		"message": "สมัครสมาชิกสำเร็จ",
	})
}

func Login(c *gin.Context) {
	c.JSON(201, gin.H{
		"data": "Login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(201, gin.H{
		"data": id,
	})
}

func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname")
	c.JSON(201, gin.H{
		"data": fullname,
	})
}
