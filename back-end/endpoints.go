package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
	config "web-service-gin/configs"
	"web-service-gin/dtos"
	"web-service-gin/models"
	"web-service-gin/services"
)

// File for handler functions, to clean up main. Production are the functions used in the program
// Debug are for backend testing purposes.

type ClassificationResponse struct {
	Message string `json:"message"`
}

// PRODUCTION FUNCTIONS

// GET handlers
// getUsers retrieves all users from the database and returns them as a JSON array.
func getUsers(c *gin.Context) {
	var users []config.User // Slice to hold the users.

	// Retrieve all users from the database.
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if retrieval fails.
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": users}) // Return the users as a JSON array.
}

// getUser retrieves a user by ID from the database and returns it as a JSON object.
func getUser(c *gin.Context) {
	id := c.Param("id") // Get the user ID from the URL parameter.

	// Retrieve the user from the database using the retrieveUser function.
	user, err := retrieveUser(id)
	if err != nil {
		// Return an error if the user is not found.
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found by ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user}) // Return the user as a JSON object.
}

// getUserByEmail retrieves a user by email from the database and returns it as a JSON object.
func getUserByEmail(c *gin.Context) {
	email := c.Query("email") // Get the email from the query parameter.

	// Retrieve the user from the database using the findUserByEmail function.
	user, err := findUserByEmail(email)
	if err != nil {
		// Return an error if the user is not found.
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found by email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user}) // Return the user as a JSON object.
}

// getClassification retrieves a user by ID and sends a POST request to a Python server with the user's image URL.
func getClassification(c *gin.Context) {
	id := c.Param("id") // Get the user ID from the URL parameter.

	// Retrieve the user from the database using the retrieveUser function.
	user, err := retrieveUser(id)
	if err != nil {
		// Return an error if the user is not found.
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found by ID"})
		return
	}

	// Call the classify function to get the classification response
	response, err := classify(user.ImgURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to classify image",
		})
		return
	}

	// Set the JSON response in the Gin context
	c.JSON(http.StatusOK, response)
}

// getClassification retrieves a user by ID and sends a POST request to a Python server with the user's image URL.
func classify(imageURL string) (ClassificationResponse, error) {

	url := "http://localhost:8080/python" // URL of the Python server.

	// Create a JSON payload with the user's image URL.
	var jsonStr = []byte(`{"img_url":"` + imageURL + `"}`)

	// Create a new POST request with the JSON payload.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{} // Create a new HTTP client.

	resp, err := client.Do(req) // Send the POST request to the Python server.
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Close the response body when done.

	fmt.Println("response Status:", resp.Status) // Print the response status.

	// Parse the JSON response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	// Access the JSON data from the response
	message, ok := response["message"].(string)
	if !ok {
		return ClassificationResponse{}, errors.New("Failed to parse JSON response")
	}

	fmt.Println("message:", message) // Print the "message" key from the response

	// Create the response struct
	returnVal := ClassificationResponse{
		Message: message,
	}
	return returnVal, nil
}

// getUserImages retrieves a user by email and lists the images in the user's directory.
func getUserImages(c *gin.Context) {
	// Declare a struct to hold the request body.
	//var req struct {
	//	Email string `json:"email" binding:"required"`
	//}
	//
	//// Bind the request body to the req struct.
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if binding fails.
	//	return
	//}
	//
	//var user config.User // Declare a variable to hold the user.
	//
	//// Retrieve the user from the database by email.
	//if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Return an error if the user is not found.
	//	return
	//}

	var user config.User // Declare a variable to hold the user.

	rawData, err := c.GetRawData() // Read the request body into a byte slice.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if reading fails.
		return
	}

	var jsonBody map[string]interface{} // Declare a variable to hold the request body as a map.

	if err := json.Unmarshal(rawData, &jsonBody); err != nil { // Unmarshal the request body into the jsonBody variable.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if unmarshaling fails.
		return
	}

	token, ok := jsonBody["token"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if token not present
		return

	}

	userID, _ := verifyToken(token.(string))
	// Retrieve the user from the database by ID.
	if err := db.First(&user, userID).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound) // Return an error if the user is not found.
		return
	}

	directory_substr := "go-cloudinary/" + user.Email // Construct the directory path.

	urls, _ := listImagesInDirectory(directory_substr) // List the images in the directory.

	c.JSON(http.StatusOK, gin.H{"urls": urls})

}

// POST handlers
// createUser creates a new user in the database and returns a JWT token.
func createUser(c *gin.Context) {
	var user config.User // Declare a variable to hold the user.

	// Bind the request body to the user variable.
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if binding fails.
		return
	}

	// Check if the email is valid using the isValidEmail function.
	if matched := isValidEmail(user.Email); !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"}) // Return an error if the email is invalid.
		return
	}

	// Check if the password is present and has a length greater than 5.
	if len(user.Password) <= 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be longer than 5 characters"}) // Return an error if the password is too short.
		return
	}

	// Check if a user with the same email already exists in the database.
	var existingUser config.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists"}) // Return an error if a user with the same email already exists.
		return
	}

	byteArray, err := HashPassword(user.Password) // Hash the user's password using the HashPassword function.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password failed to hash"}) // Return an error if hashing fails.
	}

	user.Password = string(byteArray) // Set the user's password to the hashed password.

	if err := db.Create(&user).Error; err != nil { // Create the user in the database.
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if creation fails.
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Create a new JWT token for the user.
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey)) // Sign the token with a secret key.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"}) // Return an error if signing fails.
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user}) // Return the token and user as a JSON object.
}

// verifyUser verifies a user's email and password and returns a JWT token.
func verifyUser(c *gin.Context) {
	// Declare a struct to hold the request body.
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the request body to the req struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if binding fails.
		return
	}

	var user config.User // Declare a variable to hold the user.

	// Retrieve the user from the database by email.
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Return an error if the user is not found.
		return
	}

	// Compare the hashed password in the database with the provided password.
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"}) // Return an error if the passwords don't match.
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Create a new JWT token for the user.
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey)) // Sign the token with a secret key.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"}) // Return an error if signing fails.
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user}) // Return the token and user as a JSON object.
}

// PUT handlers
// updateUser updates a user in the database and returns the updated user as a JSON object.
func updateUser(c *gin.Context) {
	var user config.User // Declare a variable to hold the user.

	rawData, err := c.GetRawData() // Read the request body into a byte slice.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if reading fails.
		return
	}

	var jsonBody map[string]interface{} // Declare a variable to hold the request body as a map.

	if err := json.Unmarshal(rawData, &jsonBody); err != nil { // Unmarshal the request body into the jsonBody variable.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if unmarshaling fails.
		return
	}

	token, ok := jsonBody["token"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if token not present
		return

	}

	userID, _ := verifyToken(token.(string))
	// Retrieve the user from the database by ID.
	if err := db.First(&user, userID).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound) // Return an error if the user is not found.
		return
	}

	if err := json.Unmarshal(rawData, &user); err != nil { // Unmarshal the request body into the user variable.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if unmarshaling fails.
		return
	}

	if _, ok := jsonBody["password"]; ok { // Check if the password is present in the request body.
		byteArray, err := HashPassword(user.Password) // Hash the new password using the HashPassword function.
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password failed to hash"}) // Return an error if hashing fails.
		}

		user.Password = string(byteArray) // Set the user's password to the hashed password.
	}

	var existingUser config.User // Declare a variable to hold an existing user.

	// Check if a user with the same email already exists in the database (excluding the current user).
	if err := db.Where("email = ? AND id != ?", user.Email, userID).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists"}) // Return an error if a user with the same email already exists.
		return
	}

	if err := db.Save(&user).Error; err != nil { // Save the updated user in the database.
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if saving fails.
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user}) // Return the updated user as a JSON object.
}

// uploadUserImageByID uploads an image for a user and updates the user's image URL in the database.
func uploadUserImageByID(c *gin.Context) {
	var user config.User // Declare a variable to hold the user.

	id := c.Param("id") // Get the user ID from the URL parameter.

	// Retrieve the user from the database by ID.
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ // Return an error if the user is not found.
			"message": "No user was found with  the given ID",
		})
		return
	}

	formfile, _, err := c.Request.FormFile("file") // Get the file from the request.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ // Return an error if no file is received.
			"message": "No file was received take an L",
		})
		return
	}

	// Upload the file using the NewMediaUpload.FileUpload function.
	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if uploading fails.
		return
	}

	user.ImgURL = uploadUrl // Set the user's image URL to the uploaded file's URL.

	db.Save(&user) // Save the updated user in the database.

	c.JSON(http.StatusOK, gin.H{"message": "image is uploaded"}) // Return a success message.
}

// uploadUserImage uploads an image for a user and updates the user's image URL in the database.
func uploadUserImage(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	userID, err := verifyToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
		return
	}

	var user config.User // Declare a variable to hold the user.

	// Retrieve the user from the database by ID.
	if err := db.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ // Return an error if the user is not found.
			"message": "No user was found with  the given ID",
		})
		return
	}

	formfile, _, err := c.Request.FormFile("file") // Get the file from the request.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ // Return an error if no file is received.
			"message": "No file was received take an L",
		})
		return
	}

	// Upload the file using the NewMediaUpload.FileUploadV2 function.
	uploadUrl, err := services.NewMediaUpload().FileUploadV2(models.File{File: formfile}, user.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if uploading fails.
		return
	}

	//// Upload the file using the NewMediaUpload.FileUpload function.
	//uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
	//if err != nil {
	//	c.AbortWithStatus(http.StatusInternalServerError) // Return an error if uploading fails.
	//	return
	//}

	user.ImgURL = uploadUrl // Set the user's image URL to the uploaded file's URL.

	db.Save(&user) // Save the updated user in the database.

	//c.JSON(http.StatusOK, gin.H{"message": "image is uploaded"}) // Return a success message.
	c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "1",
			Data:       map[string]interface{}{"data": uploadUrl},
		})
}

// uploadUserImageV2 uploads an image for a user and updates the user's image URL in the database.
func uploadUserImageV2(c *gin.Context) {
	var user config.User // Declare a variable to hold the user.

	id := c.Param("id") // Get the user ID from the URL parameter.

	// Retrieve the user from the database by ID.
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ // Return an error if the user is not found.
			"message": "No user was found with  the given ID",
		})
		return
	}

	formfile, _, err := c.Request.FormFile("file") // Get the file from the request.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ // Return an error if no file is received.
			"message": "No file was received take an L",
		})
		return
	}

	// Upload the file using the NewMediaUpload.FileUploadV2 function.
	uploadUrl, err := services.NewMediaUpload().FileUploadV2(models.File{File: formfile}, user.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if uploading fails.
		return
	}

	user.ImgURL = uploadUrl // Set the user's image URL to the uploaded file's URL.

	db.Save(&user) // Save the updated user in the database.

	c.JSON(http.StatusOK, gin.H{"message": "image is uploaded"}) // Return a success message.
}

// DELETE /users/:id endpoint
// deleteUser deletes a user from the database.
func deleteUser(c *gin.Context) {
	var user config.User // Declare a variable to hold the user.

	id := c.Param("id") // Get the user ID from the URL parameter.

	// Retrieve the user from the database by ID.
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound) // Return an error if the user is not found.
		return
	}

	db.Delete(&user) // Delete the user from the database.

	c.JSON(http.StatusOK, gin.H{"data": "user deleted"}) // Return a success message.
}

// DEBUG FUNCTIONS
// updateUserDEBUG updates a user in the database and returns the updated user as a JSON object, using ID as path param
func updateUserDEBUG(c *gin.Context) {
	var user config.User // Declare a variable to hold the user.

	id := c.Param("id") // Get the user ID from the URL parameter.

	// Retrieve the user from the database by ID.
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound) // Return an error if the user is not found.
		return
	}

	rawData, err := c.GetRawData() // Read the request body into a byte slice.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if reading fails.
		return
	}

	if err := json.Unmarshal(rawData, &user); err != nil { // Unmarshal the request body into the user variable.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if unmarshaling fails.
		return
	}

	var jsonBody map[string]interface{} // Declare a variable to hold the request body as a map.

	if err := json.Unmarshal(rawData, &jsonBody); err != nil { // Unmarshal the request body into the jsonBody variable.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return an error if unmarshaling fails.
		return
	}

	if _, ok := jsonBody["password"]; ok { // Check if the password is present in the request body.
		byteArray, err := HashPassword(user.Password) // Hash the new password using the HashPassword function.
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password failed to hash"}) // Return an error if hashing fails.
		}

		user.Password = string(byteArray) // Set the user's password to the hashed password.
	}

	var existingUser config.User // Declare a variable to hold an existing user.

	// Check if a user with the same email already exists in the database (excluding the current user).
	if err := db.Where("email = ? AND id != ?", user.Email, id).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists"}) // Return an error if a user with the same email already exists.
		return
	}

	if err := db.Save(&user).Error; err != nil { // Save the updated user in the database.
		c.AbortWithStatus(http.StatusInternalServerError) // Return an error if saving fails.
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user}) // Return the updated user as a JSON object.
}
