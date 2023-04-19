package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/dtos"
	"web-service-gin/models"
	"web-service-gin/services"
)

type ClassificationResponse struct {
	Message string `json:"message"`
}

func FileUploadWithClassification() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upload
		formfile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		// Call the classify function to get the classification response
		response, err := classify(uploadUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to classify image",
			})
			return
		}

		// Set the JSON response in the Gin context
		c.JSON(http.StatusOK, response)

	}
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

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upload
		formfile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}

func RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url models.Url

		//validate the request body
		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				dtos.MediaDto{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
