package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var (
	DB *gorm.DB
	err error
)

type User struct {
	gorm.Model
	Type     string `gorm:"default:0"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Wishes	 []Wish
	Username string `gorm:"type:unique_index;not null"`
	Password string `gorm:"not null"`
}

type Wish struct {
	gorm.Model
	UserID   int
	User     User
	Name     string
	Email    string
	Message  string
	Number   string
	DateTime time.Time
}

func NewDatabase() {
	DB, err = gorm.Open("sqlite", "wisbday.sqlite")
	if err != nil {
		panic("Could't connect to database")
	}

	defer func() {
		if err := DB.Close(); err != nil {
			panic("Could not close the database connection")
		}
	}()

	DB.AutoMigrate(&User{}, &Wish{})

	fmt.Println("Connected to database successfully")
}
