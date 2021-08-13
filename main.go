package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofirebase/api"
	"gofirebase/config"
)

func main() {

	// initialize new gin engine (for server)
	r := gin.Default()
	fmt.Println("14")
	// create/configure database instance
	db := config.CreateDatabase()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	// set db & firebase auth to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	// routes definition for finding and creating artists
	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)

	// start the server
	r.Run(":5000")

}
