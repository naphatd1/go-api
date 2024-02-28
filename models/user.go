package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	IsAdmin   bool      `json:"is_admin" gorm:"type:bool;default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = HashPassword(user.Password)
	return nil
}

func HashPassword(password string) string {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		panic(err) // 💥
	}
	return string(encoded)
}
