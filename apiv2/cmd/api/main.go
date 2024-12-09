package main

import (
	"apiv2/pkg/db"
	"apiv2/pkg/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT string = ":8080"

func main() {
	// this is a pointer
	server := gin.Default()
	db.InitDB()
	routes.RegisterRoutes(server)
	fmt.Printf("server listening on http://localhost%v\n", PORT)

	err := server.Run(PORT)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
	}
}
