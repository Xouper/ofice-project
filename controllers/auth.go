package controllers

import (
	//"fmt"
	"learning/db"

	"github.com/gin-gonic/gin"
)

func CreatUser(u *db.User) {
	if u.Age > 18 {
		u.Adult = true
	}
	db.Users = append(db.Users, *u)
}

func SaveUser(c *gin.Context) {
	var body db.User
	c.BindJSON(&body)
	for _, value := range db.Users {
		if body.Username == value.Username && body.Gender == value.Gender{
			return
		}
	}
	CreatUser(&body)
}

func SaveAsAdult(users []db.User) []db.User {
	var Adults []db.User
	for i := 0; i < len(users); i++ {
		if users[i].Adult {
			Adults = append(Adults, users[i])
		}
	}
	return Adults
}

func ShowAllUsers(c *gin.Context) {
	c.JSON(200, db.Users)
}

func ShowAdults(c *gin.Context) {
	Adults := SaveAsAdult(db.Users)
	c.JSON(200, Adults)
}
