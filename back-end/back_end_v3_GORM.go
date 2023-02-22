package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	//"github.com/gin-contrib/limit"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"regexp"
	config "web-service-gin/configs"
	"web-service-gin/controllers"
	"web-service-gin/models"
	"web-service-gin/services"
)

// GET handlers
func getUsers(c *gin.Context) {
	var users []config.User

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": users})
} //GET all the users, /users

func getUser(c *gin.Context) {
	var user config.User
	id := c.Param("id")
	fmt.Println(id)
	if err := db.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found by id"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
} //GET a specific user, /user/:id

func getUserByEmail(c *gin.Context) {
	email := c.Query("email")
	fmt.Println("This is the email: ")
	fmt.Println(email)
	user, err := findUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found by email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST handlers
func createUser(c *gin.Context) {
	var user config.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if matched := isValidEmail(user.Email); !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Check if the email already exists in the database
	var existingUser config.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
} //POST (create) a new user on SQLite database, /users

func verifyUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user config.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	fmt.Println(req.Email)
	fmt.Println(req.Password)

	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err := db.Where("password = ?", req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//func verifyUser(c *gin.Context) {
//	email := c.Param("email")
//	password := c.Query("password")
//
//	var user config.User
//	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		return
//	}
//
//	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}

// PUT /users/:id endpoint
func updateUser(c *gin.Context) {
	var user config.User
	id := c.Param("id")

	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email already exists in the database
	var existingUser config.User

	if err := db.Where("email = ? AND id != ?", user.Email, id).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists"})
		return
	}

	if err := db.Save(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
} //PUT updated information for a user, /users/:id

func uploadUserImage(c *gin.Context) {
	var user config.User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No user was found with  the given ID",
		})
		return
	}

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
	var user config.User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.First(&user, id)
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "user deleted"})
}

//Helper functions

func findUserByEmail(email string) (*config.User, error) {
	var user config.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func isValidEmail(email string) bool {
	matched, err := regexp.MatchString(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, email)
	if err != nil {
		return false
	}
	return matched
}

var db *gorm.DB

func main() {

	//Connect to the SQLite database
	db_temp, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//initialize package level db variable
	db = db_temp

	// Automatically create the "users" table based on the User struct
	db.AutoMigrate(&config.User{})

	// Create a new Gin router
	r := gin.Default()

	// Set a rate limit of 10 requests per second with a burst of 5.
	//r.Use(limit.NewRateLimiter(10, 5))

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200", "http://localhost:8080", "http://localhost:8081"}
	r.Use(cors.New(config))

	// Define the routes
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.GET("/users/email", getUserByEmail)

	r.POST("/users", createUser)
	r.POST("/users/login", verifyUser)

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
