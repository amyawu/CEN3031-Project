package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
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
