package db

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SaveNewUserToFile(u *User) {
	file, err := os.OpenFile("db/users.csv", os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	s := []string{
		Converter(u.ID),
		u.Username,
		u.Password,
		Converter(u.Height),
		u.Gender,
		Converter(u.Age),
	}
	w := csv.NewWriter(file)
	w.Write(s)

	w.Flush()
	file.Close()
}

func Converter(vl interface{}) string {
	return fmt.Sprintf("%v", vl)
}
