package misc_files

//import (
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/sqlite"
//)
//
//// User struct
//type User struct {
//	gorm.Model
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}
//
//func main() {
//	// Connect to the SQLite database
//	db, err := gorm.Open("sqlite3", "test.db")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Automatically create the "users" table based on the User struct
//	db.AutoMigrate(&User{})
//
//	// Create a new Gin router
//	r := gin.Default()
//
//	// Define the GET /users endpoint
//	r.GET("/users", func(c *gin.Context) {
//		var users []User
//		db.Find(&users)
//		c.JSON(http.StatusOK, gin.H{"data": users})
//	})
//
//	// Define the POST /users endpoint
//	r.POST("/users", func(c *gin.Context) {
//		var user User
//		if err := c.ShouldBindJSON(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		db.Create(&user)
//		c.JSON(http.StatusOK, gin.H{"data": user})
//	})
//
//	// Define the PUT /users/:id endpoint
//	r.PUT("/users/:id", func(c *gin.Context) {
//		var user User
//		id := c.Param("id")
//		db.First(&user, id)
//		if err := c.ShouldBindJSON(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		db.Save(&user)
//		c.JSON(http.StatusOK, gin.H{"data": user})
//	})
//
//	// Define the DELETE /users/:id endpoint
//	r.DELETE("/users/:id", func(c *gin.Context) {
//		var user User
//		id := c.Param("id")
//		db.First(&user, id)
//		db.Delete(&user)
//		c.JSON(http.StatusOK, gin.H{"data": "user deleted"})
//	})
//
//	// Start the server
//	fmt.Println("Server is running on port 8000...")
//	r.Run(":8000")
//}
