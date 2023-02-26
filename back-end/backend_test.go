package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
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

	// Test if the router is not nil
	if r == nil {
		t.Errorf("Router is nil")
	}
}

func TestCreateUser(t *testing.T) {
	// Create a new Gin router
	r := SetUpRouter()

	// Define the routes
	r.POST("/users", createUser)

	// Create a new user to add to the database
	newUser := config.User{
		Name:  "John Smith",
		Email: "john.smith@example.com",
	}

	// Convert the user struct to JSON
	jsonData, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("error marshaling JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "localhost:8080/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	//// Create a new HTTP response recorder
	//w := httptest.NewRecorder()
	//
	//// Perform the request
	//r.ServeHTTP(w, req)
	//
	//// Check the response status code
	//if w.Code != http.StatusCreated {
	//	t.Errorf("expected status code %d but got %d", http.StatusCreated, w.Code)
	//}

	//// Check that the response body contains the created user
	//var createdUser config.User
	//err = json.Unmarshal(w.Body.Bytes(), &createdUser)
	//if err != nil {
	//	t.Errorf("error unmarshaling JSON: %v", err)
	//}
	//if createdUser.Name != newUser.Name {
	//	t.Errorf("expected user name %s but got %s", newUser.Name, createdUser.Name)
	//}
	//if createdUser.Email != newUser.Email {
	//	t.Errorf("expected user email %s but got %s", newUser.Email, createdUser.Email)
	//}
}
