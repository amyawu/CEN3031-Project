package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	config "web-service-gin/configs"
	"web-service-gin/controllers"
)

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

var jwtKey = []byte("ZkYcqWwhjK/TSFMY2eL21mZADY9x0w+UAqF4UwIRaAY=")

func main() {
	/*
		opt := option.WithCredentialsFile("./private/melagoomba-firebase.json") // Firebase service key
		firebase_config := &firebase.Config{
			DatabaseURL: "https://melagoomba-default-rtdb.firebaseio.com/", // Firebase database URL
		}

		app, err := firebase.NewApp(context.Background(), firebase_config, opt)
		if err != nil {
			log.Fatalf("Failed to intialize Firebase app: %v", err)
		}
		client, err := app.Database(context.Background())
		if err != nil {
			log.Fatalf("Failed to get Firebase database client: %v", err)
		}

		db, err := gorm.Open("sqlite3", .Ref("").URL.String()+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalf("Failed to open database connection: %v", err)
		}
		defer db.Close()
	*/

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

	r.PUT("/users/:id", updateUserDEBUG)
	r.PUT("/users/profile", updateUser)

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
