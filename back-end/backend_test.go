package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	config "web-service-gin/configs"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRouter(t *testing.T) {
	// Create a new router
	r := SetUpRouter()
	r.GET("/users/:id", getUser)

	// Test if the router is not nil
	if r == nil {
		t.Errorf("Router is nil")
	}
}

func TestGetUser(t *testing.T) {
	// Create a new test router
	r := gin.New()
	//Test User to grab
	//expectedUser, _ := retrieveUser("1")

	// Define the route and handler function
	r.GET("/users/:id", getUser)
	// Create a new HTTP request to the test router with the user's ID
	req, _ := http.NewRequest("GET", "/users/"+fmt.Sprint(1), nil)
	// Create a new HTTP recorder to capture the response
	w := httptest.NewRecorder()
	// Serve the HTTP request to the test router
	r.ServeHTTP(w, req)

	// Check if the response code is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

}

func TestVerifyUser(t *testing.T) {
	// Create a new test router
	r := gin.New()
	//Test User to grab
	//user := config.User{Name: "Emily Bronte", Email: "ebronte@gmail.com", Password: "password"}

	// Define the route and handler function
	r.POST("/users/login", verifyUser)
	// Create a new HTTP request to the test router with the user's ID
	req, _ := http.NewRequest("GET", "/users/"+fmt.Sprint(1), nil)

	// Create a new HTTP recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the HTTP request to the test router
	r.ServeHTTP(w, req)

	// Check if the response code is 200 OK
	//if w.Code != http.StatusOK {
	//	t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	//}

}

func TestGetUserByEmail(t *testing.T) {
	// Create a new test router
	r := gin.New()
	//Test User to grab
	user := config.User{Name: "Emily Bronte", Email: "ebronte@gmail.com", Password: "password"}
	//db.Create(&user)

	// Define the route and handler function
	r.GET("/users/email", getUserByEmail)
	fmt.Println(user.ID)
	// Create a new HTTP request to the test router with the user's ID
	req, _ := http.NewRequest("GET", "/users/email"+"?email=David1234@gmail.com", nil)

	// Create a new HTTP recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the HTTP request to the test router
	r.ServeHTTP(w, req)

	// Check if the response code is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

}

func TestCreateUser(t *testing.T) {
	// Create a new Gin router
	r := SetUpRouter()

	// Define the routes
	r.POST("/users", createUser)

	// Create a new user to add to the database
	newUser := config.User{
		Name:     "John Smith",
		Email:    "john.smith@example.com",
		Password: "password",
	}

	// Convert the user struct to JSON
	jsonData, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("error marshaling JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

}

// Test to make sure it hashes correct, passes if there's no errors and a hash exists
func TestHashPassword(t *testing.T) {
	password := "password123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if len(hash) == 0 {
		t.Errorf("Hashed password is empty.")
	}
}

// Test to see if log in is possible, will pass if the two passwords match
func TestUserLogin(t *testing.T) {
	password := "password123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		t.Errorf("Passwwords do not match: %v", err)
	}
}

func TestDatabaseConnection(t *testing.T) {
	// THIS TEST WILL ALWAYS FAIL BECAUSE MICROSOFT AZURE FAILED US D:

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")

	db, err := gorm.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		username,
		password,
		host,
		port,
		dbname,
	))
	if err != nil {
		t.Errorf("Error connecting to Azure SQL database: %s", err.Error())
	}

	defer db.Close()

	if err := db.DB().Ping(); err != nil {
		t.Errorf("Error pinging Azure SQL database: %s", err.Error())
	}
}