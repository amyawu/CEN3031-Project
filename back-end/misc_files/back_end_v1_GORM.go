package misc_files

//import (
//	"fmt"
//	"log"
//
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/sqlite"
//)
//
//type User struct {
//	gorm.Model
//	Name string
//}
//
//func main() {
//	// Connect to the SQLite database
//	db, err := gorm.Open("sqlite3", "test.db")
//	if err != nil {
//		log.Fatal("Error connecting to database:", err)
//	}
//	defer db.Close()
//
//	// Automatically create the "users" table based on the User struct
//	db.AutoMigrate(&User{})
//
//	// Insert a user
//	user := User{Name: "John Doe"}
//	if err := db.Create(&user).Error; err != nil {
//		log.Fatal("Error inserting user:", err)
//	}
//	fmt.Println("User inserted with ID", user.ID)
//
//	// Query the users
//	var users []User
//	if err := db.Find(&users).Error; err != nil {
//		log.Fatal("Error querying users:", err)
//	}
//	//fmt.Println("Users:", users)
//
//	for _, user := range users {
//		fmt.Println(user.Name)
//	}
//}
