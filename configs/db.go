package configs

import (
	"fmt"
	"os"

	"github.com/naphat/gob-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASE_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ไม่สามารถติดต่อกับ Database Server ได้")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("ติดต่อฐานข้อมูลสำเร็จ")

	// Migration
	// db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	DB = db

}
