package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"web-service-gin/controllers"
	"web-service-gin/helper"
	"web-service-gin/models"
	"web-service-gin/services"
)

// User struct
type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	ImgURL string `json:"imgurl"`
}

// GET handlers
func getUsers(c *gin.Context) {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": users})
} //GET all the users, /users

func getUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		//fmt.Println("The user was not found??")
		//c.AbortWithStatus(http.StatusNotFound)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
} //GET a specific user, /user/:id

// POST handlers
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
} //POST (create) a new user on SQLite database, /users

func createUserWithImage(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error is happening here": err.Error()})
		return
	}

	// Get image file from the request
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	// Upload image to Cloudinary
	imageUrl, err := helper.ImageUploadHelper(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save user data and image URL in the database
	user.ImgURL = imageUrl
	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
} //POST (create) a new user, with a corresponding image in Cloudinary, /users/image

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
} //PUT updated information for a user, /users/:id

func uploadUserImage(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No user was found with  the given ID",
		})
		return
	}
	//controllers.FileUpload()

	formfile, _, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file was received take an L",
		})
		return
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user.ImgURL = uploadUrl
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "image is uploaded"})
} //PUT a new image for an existing user, /users/:id/image

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
	r.POST("/users/image", createUserWithImage) //May not be necessary

	r.PUT("/users/:id", updateUser)
	r.PUT("/users/:id/image", uploadUserImage)

	r.DELETE("/users/:id", deleteUser)

	//Image upload
	r.POST("/file", controllers.FileUpload())
	r.POST("/remote", controllers.RemoteUpload())

	// Start the server
	fmt.Println("Server is running on port 8000...")
	r.Run(":8000")
}
