package main

import (
	"learning/controllers"

	"learning/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.POST("/auth", controllers.SaveUser)
	r.GET("/users", controllers.ShowAllUsers)
	r.GET("/", controllers.CreatQrc)
	r.POST("/users/:id/image/", controllers.AddUserImage)
	r.GET("/users/:id/image/", controllers.ShowUserImage)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//controllers.ParseHabr()
}
