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
		if body.Username == value.Username && body.Gender == value.Gender {
			return
		}
	}
	CreatUser(&body)
}

func ReturnAdults(users []db.User) []db.User {
	var Adults []db.User
	for i := 0; i < len(users); i++ {
		if users[i].Adult {
			Adults = append(Adults, users[i])
		}
	}
	return Adults
}

func FilterGender(users []db.User, gender string) []db.User {
	var FilteredUsers []db.User
	for i := 0; i < len(users); i++ {
		if users[i].Gender  == gender {
			FilteredUsers = append(FilteredUsers, users[i])
		}
	}
	return FilteredUsers
}

func ShowAllUsers(c *gin.Context) {
	users := db.Users

	QueryAdult := c.DefaultQuery("adult", "default")

	if QueryAdult != "default" {
		users = ReturnAdults(users)
	}

	QueryGender := c.Query("gender")
	
	if QueryGender != ""{
		users = FilterGender(users, QueryGender)
	}
	c.JSON(200, users)
}
	