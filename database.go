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
	Type     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Wishes	 []Wish
	Username string `gorm:"type:unique_index"`
	Password string
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

	defer DB.Close()
	DB.AutoMigrate(&User{}, &Wish{})

	fmt.Println("Connected to database successfully")
}
