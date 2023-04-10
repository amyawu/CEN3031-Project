package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	config "web-service-gin/configs"
)

// File for helper functions, will be used by endpoints.go's handler functions which are called in main

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

	// Create a slice of URLs
	var urls []string
	for _, resource := range resources.Assets {
		fmt.Println(resource.SecureURL)
		urls = append(urls, resource.SecureURL)
	}

	return urls, nil
}

func genKey() {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(key))
}

func verifyToken(tokenString string) (int, error) {
	// Parse the token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return jwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return int(token.Claims.(jwt.MapClaims)["sub"].(float64)), nil
}
