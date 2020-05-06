package models

import (
	"Auth/Cms_Authentication_Jwt/api/security"
	"log"
	"time"
)

// User struct
type User struct {
	Id            uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Firstname     string     `gorm:"size:20;not null;unique_index" json:"firstname,omitempty"`
	Lastname      string     `gorm:"size:20;not null;unique_index" json:"lastname,omitempty"`
	Email         string     `gorm:"size:35;not null;unique_index" json:"email,omitempty"`
	Password      string     `gorm:"size:60;not null" json:"password,omitempty"`
	AccountStatus bool       `gorm:"not null; default:true" json:"accountstatus"`
	Address       string     `gorm:"size:60;not null" json:"address,omitempty"`
	Organisation  string     `gorm:"size:60;not null" json:"organisation,omitempty"`
	PhoneNumber   string     `gorm:"size:15;not null" json:"phonenumber"`
	Photo         string     `gorm:"size:60;not null" json:"photo,omitempty"`
	PhotoLocation string     `gorm:"size:60;not null" json:"photolocation,omitempty"`
	Role          string     `gorm:"size:60;not null" json:"role,omitempty"`
	CreatedAt     *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func CreateUser(user User) (interface{}, error) {
	db := Connect()
	defer db.Close()
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashedPassword)
	rs := db.Create(&user)
	return rs.Value, rs.Error
}

func GetUsers() []User {
	db := Connect()
	defer db.Close()
	var users []User
	db.Order("id asc").Find(&users)
	return users
}

func GetUserByEmail(email string) User {
	db := Connect()
	defer db.Close()
	var user User
	db.Where("email = ?", email).Find(&user)
	return user
}
