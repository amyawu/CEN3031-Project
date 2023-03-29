package main

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os/exec"
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
	id := c.Param("id")

	user, err := retrieveUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found by ID"})
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

func getClassification(c *gin.Context) {

	id := c.Param("id")
	user, err := retrieveUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found by ID"})
		return
	}
	//Requires Conda environment names 'Tensorflow' with Tensorflow v.X.X.X installed
	cmd := exec.Command("conda", "activate", "Tensorflow", "&&", "python", "./python_scripts/sample_network.py", user.ImgURL)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
	}
	fmt.Println(string(out))
	c.IndentedJSON(http.StatusOK, gin.H{"data": string(out)})

}

func getUserImages(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
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

	directory_substr := "go-cloudinary/" + user.Email
	listImagesInDirectory(directory_substr)
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

	byteArray, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password failed to hash"})
	}

	// Saved the hashed byte array password as a string, may change later.
	user.Password = string(byteArray)

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
	fmt.Println(user.Password) //OG
	fmt.Println(req.Email)
	fmt.Println(req.Password) //hashed

	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err := db.Where("password = ?", req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, user)
}

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

	var req struct {
		Email string `json:"email" binding:"required"`
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

func uploadUserImageByID(c *gin.Context) {

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

func uploadUserImageV2(c *gin.Context) {

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

	uploadUrl, err := services.NewMediaUpload().FileUploadV2(models.File{File: formfile}, user.Email)
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
	//db, _ = openDB()

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

// HashPassword returns a byte slice containing the bcrypt hash of the password at the given cost.
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash compares a bcrypt hashed password with the inputted password. Returns nil on success, or an error on failure.
func CheckPasswordHash(password, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

func retrieveUser(id string) (config.User, error) {
	var user config.User

	if err := db.First(&user, id).Error; err != nil {
		return config.User{}, err
	}
	return user, nil
}

func listImagesInDirectory(dirName string) ([]string, error) {

	cld, _ := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())

	// Set context
	ctx := context.Background()

	fmt.Println("All folder names:")
	resp, err := cld.Admin.RootFolders(ctx, admin.RootFoldersParams{})
	for _, resource := range resp.Folders {
		fmt.Println(resource.Name)
		fmt.Println(resource.Path)

	}

	// List resources in folder
	resources, err := cld.Admin.Assets(ctx, admin.AssetsParams{Prefix: dirName, DeliveryType: "upload"})
	if err != nil {
		return nil, err
	}

	// List resources in folder
	//resources, err := cld.Admin.ResourcesByContext(ctx, admin.ResourcesByContextParams{Prefix: folderPath})

	// Print URLs of all images within a certain directory
	for _, resource := range resources.Assets {
		fmt.Println(resource.SecureURL)
	}

	return nil, nil
}

func openDB() *gorm.DB {
	//Connect to the SQLite database

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Automatically create the "users" table based on the User struct
	if err := db.AutoMigrate(&config.User{}).Error; err != nil {
		db.Close()
		return nil
	}

	//-----------------------
	//dsn := "manuel.cortes:cen3031melanoma@tcp(mysql.cise.ufl.edu:3306)/Users" //?charset=utf8mb4
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//db.AutoMigrate(&config.User{})
	//----------------------

	return db
}

var db *gorm.DB = openDB()

func main() {

	//Database is being initialized globally
	var err error
	if db == nil {
		log.Fatal(err)
	}

	// Create a new Gin router
	r := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200", "http://localhost:8080", "http://localhost:8081"}
	r.Use(cors.New(config))

	// Define the routes
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.GET("/users/email", getUserByEmail)
	r.GET("/users/:id/classify", getClassification)
	r.GET("/users/all_images", getUserImages)

	r.POST("/users", createUser)
	r.POST("/users/login", verifyUser)

	r.PUT("/users/:id", updateUser)
	r.PUT("/users/:id/image", uploadUserImageByID)
	r.PUT("/users/image", uploadUserImage)
	r.PUT("/users/:id/imageV2", uploadUserImageV2)

	r.DELETE("/users/:id", deleteUser)

	//Image upload
	r.POST("/file", controllers.FileUpload()) // localhost:8000/file  --> Cloudinary
	r.POST("/remote", controllers.RemoteUpload())

	// Start the server
	fmt.Println("Server is running on port 8000...")
	r.Run(":8000")
}
