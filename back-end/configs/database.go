package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User struct
type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ImgURL    string `json:"imgurl"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Ethnicity int    `json:"ethnicity"`
}

func CreateDatabase() *gorm.DB {
	// Create db instance with gorm
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	// migrate our model for artist
	db.AutoMigrate(&User{})
	return db
}
