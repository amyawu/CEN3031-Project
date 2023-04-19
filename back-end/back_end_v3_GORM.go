package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	config "web-service-gin/configs"
	"web-service-gin/controllers"
)

// Microsoft Azure SQL - Go Connection Code (from documentation)
// var db *sql.DB
var server = os.Getenv("DB_SERVER")
var user = os.Getenv("DB_USER")
var port = os.Getenv("DB_PORT")
var password = os.Getenv("DB_PASSWORD") // Not good practice to have password saved out in the open, need to define it in the shell as an environment variable
var database = os.Getenv("DB_DATABASE")

func openDB() *gorm.DB {

	////Connect to the SQLite database - MANUEL
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Automatically create the "users" table based on the User struct
	if err := db.AutoMigrate(&config.User{}).Error; err != nil {
		db.Close()
		return nil
	}

	//Connect to the SQLite database - GAVIN
	//var errenv error
	//errenv = godotenv.Load(".env")
	//if errenv != nil {
	//	log.Fatalf("Error loading .env.local file: %v", errenv)
	//}
	//
	//// Azure-Go Connection Code (cont'd)
	////connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	////connString := fmt.Sprintf("sqlserver://%s:%s@%s.database.windows.net:%d?database=%s", user, password, server, port, database)
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=true;"+
	//	"connection timeout=30;", server, user, password, database, port)
	//
	//db, err := gorm.Open("sqlserver", connString)
	//if err != nil {
	//	log.Fatal("Error creating connection pool: ", err.Error())
	//}
	//
	//// Automatically create the "users" table based on the User struct
	//if err := db.AutoMigrate(&config.User{}).Error; err != nil {
	//	db.Close()
	//	return nil
	//}
	//
	return db
}

var db *gorm.DB = openDB()

var jwtKey = []byte("ZkYcqWwhjK/TSFMY2eL21mZADY9x0w+UAqF4UwIRaAY=")

func main() {
	//var errenv error
	//errenv = godotenv.Load(".env")
	//if errenv != nil {
	//	log.Fatalf("Error loading .env.local file: %v", errenv)
	//}
	//
	//// Azure-Go Connection Code (cont'd)
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	//var err error
	//db, err = sql.Open("sqlserver", connString)
	//if err != nil {
	//	log.Fatal("Error creating connection pool: ", err.Error())
	//}
	//ctx := context.Background()
	//err = db.PingContext(ctx)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Printf("Connected!")

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
	config.AllowHeaders = []string{"Authorization"}
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
	r.POST("/file/classify", controllers.FileUploadWithClassification()) // localhost:8000/file  --> Cloudinary

	r.POST("/file", controllers.FileUpload()) // localhost:8000/file  --> Cloudinary
	r.POST("/remote", controllers.RemoteUpload())

	// Start the server
	fmt.Println("Server is running on port 8000...")
	r.Run(":8000")
}
