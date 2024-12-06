package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}
type ToDo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"status"`
	CreatedAt string
	User
}

var currentTime = time.Now().Format("12th August 2023 12:02PM")

var todos = []ToDo{

	{ID: "1", Item: "Testing", Completed: false, CreatedAt: currentTime,
		User: User{
			ID:        "1",
			FirstName: "Jowie",
			Email:     "testme@gmail.com",
			CreatedAt: currentTime,
		},
	}}

func main() {
	// logger and recovery middleware, avoids crashes
	server := gin.Default()
	// handlers are functions when url is called
	server.GET("/todo", ToDoGetter)
	err := server.Run(":8080")
	log.Println("running on http://localhost:8080")
	if err != nil {
		return
	}
}

// gets all Todos in the array
func ToDoGetter(c *gin.Context) {
	// convert to JSON first
	c.IndentedJSON(http.StatusOK, todos)
}

// context is usually a pointer to gin context if registered as a handler
//func getEvents(context *gin.Context) {
//	// status code, data
//	context.JSON(http.StatusOK, gin.H{"Hello": "World"})
//}
