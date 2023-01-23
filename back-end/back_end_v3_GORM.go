package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"web-service-gin/controllers"
)

// User struct
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GET /users endpoint
func getUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func getUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users endpoint
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PUT /users/:id endpoint
func updateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id endpoint
func deleteUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.First(&user, id)
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "user deleted"})
}

var db *gorm.DB

func main() {
	// Connect to the SQLite database
	db_temp, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//initialize package level db variable
	db = db_temp

	// Automatically create the "users" table based on the User struct
	db.AutoMigrate(&User{})

	// Create a new Gin router
	r := gin.Default()

	// Define the routes
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	//Image upload
	r.POST("/file", controllers.FileUpload())
	r.POST("/remote", controllers.RemoteUpload())

	// Start the server
	fmt.Println("Server is running on port 8000...")
	r.Run(":8000")
}
