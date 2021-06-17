package db

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Height   float64 `json:"height"`
	Gender   string  `json:"gender"`
	Age      int     `json:"age"`
	Adult    bool
}

var Users []User
