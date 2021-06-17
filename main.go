package main

import (
	"learning/controllers"

	// "github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()

	// r.GET("/parse", controllers.SaveUser.ParseHabr)
	// r.POST("/auth", controllers.SaveUser)
	// r.GET("/users", controllers.ShowAllUsers)
	// r.GET("/adults", controllers.ShowAdults)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	controllers.ParseHabr()
}