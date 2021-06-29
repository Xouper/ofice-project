package db

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Height   float64 `json:"height"`
	Gender   string  `json:"gender"`
	Age      int     `json:"age"`
	Adult    bool    `json:"adult"`
}

var Users []User

func AddId(user *User) {
	user.ID = len(Users) + 1
}
