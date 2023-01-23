package misc_files

//
//
//import (
//	"fmt"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//// user represents data about a record user.
//type user struct {
//	ID     string `json:"id"`
//	Name   string `json:"name"`
//	ImgURL string `json:"imgurl"`
//	Age    int    `json:"age"`
//}
//
//// getusers responds with the list of all users as JSON.
//func getUsers(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, users)
//}
//
//// postusers adds a user from JSON received in the request body.
//func postUsers(c *gin.Context) {
//	var newUser user
//	print(newUser.Name)
//	print(newUser.ImgURL)
//	print(newUser.Age)
//	// Call BindJSON to bind the received JSON to
//	// newuser.
//	if err := c.BindJSON(&newUser); err != nil {
//		fmt.Println("Post command FAILED")
//		return
//	}
//	// Add the new user to the slice.
//	users = append(users, newUser)
//	c.IndentedJSON(http.StatusCreated, newUser)
//}
//
//// getUserByID locates the user whose ID value matches the id
//// parameter sent by the client, then returns that user as a response.
//
//func getUserByID(c *gin.Context) {
//	id := c.Param("id")
//
//	// Loop over the list of users, looking for
//	// an user whose ID value matches the parameter.
//	for _, a := range users {
//		if a.ID == id {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
//}
//
//var users = []user{
//	{
//		ID:     "1",
//		Name:   "Adam",
//		ImgURL: "sample1",
//		Age:    23,
//	},
//	{
//		ID:     "2",
//		Name:   "Blake",
//		ImgURL: "sample2",
//		Age:    25,
//	},
//	{
//		ID:     "3",
//		Name:   "Connor",
//		ImgURL: "sample3",
//		Age:    27,
//	},
//}
//
//func main() {
//	router := gin.Default()
//	router.GET("/users", getUsers)
//	router.GET("/users/:id", getUserByID)
//	router.POST("/users", postUsers)
//
//	router.Run("localhost:8080")
//}
