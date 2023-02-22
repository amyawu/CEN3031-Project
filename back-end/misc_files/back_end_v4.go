package misc_files

//import (
//	"context"
//	"fmt"
//	"github.com/gin-contrib/cors"
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/sqlite"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"golang.org/x/crypto/bcrypt"
//	"gorm.io/driver/mongo"
//	"gorm.io/gorm"
//	"log"
//	"net/http"
//	config "web-service-gin/configs"
//	"web-service-gin/controllers"
//	"web-service-gin/models"
//	"web-service-gin/services"
//)
//
//type User struct {
//	gorm.Model
//	Name     string `json:"name"`
//	Email    string `json:"email"`
//	Password string `json:"password"`
//	ImgURL   string `json:"imgurl"`
//	Gender   string `json:"gender"`
//	Age      int    `json:"age"`
//}
//
//func getUsers(c *gin.Context) {
//	var users []User
//	db.Find(&users)
//	c.JSON(200, gin.H{"data": users})
//}
//
//func getUser(c *gin.Context) {
//	var user User
//	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
//		c.JSON(404, gin.H{"error": "User not found"})
//		return
//	}
//	c.JSON(200, gin.H{"data": user})
//}
//
//func getUserByEmail(c *gin.Context) {
//	// Get the email parameter from the request URL
//	email := c.Param("email")
//
//	// Look up the user by email in the users collection
//	var user User
//	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		return
//	}
//
//	// Return the user as JSON
//	c.JSON(http.StatusOK, user)
//}
//
//func verifyUser(c *gin.Context) {
//	email := c.Param("email")
//	password := c.Query("password")
//
//	var user User
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
//
//func createUser(c *gin.Context) {
//	var user User
//	if err := c.BindJSON(&user); err != nil {
//		c.JSON(400, gin.H{"error": "Bad request"})
//		return
//	}
//	db.Create(&user)
//	c.JSON(201, gin.H{"data": user})
//}
//
//func updateUser(c *gin.Context) {
//	var user User
//	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
//		c.JSON(404, gin.H{"error": "User not found"})
//		return
//	}
//	if err := c.BindJSON(&user); err != nil {
//		c.JSON(400, gin.H{"error": "Bad request"})
//		return
//	}
//	db.Save(&user)
//	c.JSON(200, gin.H{"data": user})
//}
//
//func deleteUser(c *gin.Context) {
//	var user User
//	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
//		c.JSON(404, gin.H{"error": "User not found"})
//		return
//	}
//	db.Delete(&user)
//	c.JSON(200, gin.H{"data": true})
//}
//
//func uploadUserImage(c *gin.Context) {
//	var user config.User
//	id := c.Param("id")
//	if err := db.First(&user, id).Error; err != nil {
//		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
//			"message": "No user was found with  the given ID",
//		})
//		return
//	}
//	//controllers.FileUpload()
//
//	formfile, _, err := c.Request.FormFile("file")
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//			"message": "No file was received take an L",
//		})
//		return
//	}
//
//	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	user.ImgURL = uploadUrl
//	db.Save(&user)
//	c.JSON(http.StatusOK, gin.H{"message": "image is uploaded"})
//} //PUT a new image for an existing user, /users/:id/image
//
//var db *gorm.DB
//
//func main() {
//
//	// Set up MongoDB connection options
//	clientOptions := options.Client().ApplyURI("mongodb+srv://4n1m4t10n:porfidio@cluster0.2oravn6.mongodb.net/?retryWrites=true&w=majority")
//
//	// Connect to MongoDB using the GORM driver
//	client, err := mongo.NewClient(clientOptions)
//	if err != nil {
//		panic(err)
//	}
//	err = client.Connect(context.Background())
//	if err != nil {
//		panic(err)
//	}
//
//	// Create a new GORM DB instance
//	db, err := gorm.Open(mongo.NewClient(clientOptions), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//
//	// Connect to the MongoDB database
//	db_temp, err := gorm.Open("mongo", "mongodb://localhost:27017/test")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db_temp.Close()
//
//	// Initialize package level db variable
//	db = db_temp
//
//	// Automatically create the "users" collection based on the User struct
//	db.AutoMigrate(&User{})
//
//	// Create a new Gin router
//	r := gin.Default()
//
//	// Enable CORS
//	config := cors.DefaultConfig()
//	config.AllowOrigins = []string{"http://localhost:4200", "http://localhost:8080", "http://localhost:8081"}
//	r.Use(cors.New(config))
//
//	// Define the routes
//	r.GET("/users", getUsers)
//	r.GET("/users/:id", getUser)
//	r.GET("/users/:email", getUserByEmail)
//
//	r.POST("/users", createUser)
//
//	r.PUT("/users/:id", updateUser)
//	r.PUT("/users/:id/image", uploadUserImage)
//
//	r.DELETE("/users/:id", deleteUser)
//
//	//Image upload
//	r.POST("/file", controllers.FileUpload())
//	r.POST("/remote", controllers.RemoteUpload())
//
//	// Start the server
//	fmt.Println("Server is running on port 8000...")
//	r.Run(":8000")
//}
