package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func InitDB() {
	file, err := os.Open("db/users.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		var user User
		user.ID, e = strconv.Atoi(record[0])
		user.Username = record[1]
		user.Password = record[2]
		user.Height = ParseFloat(record[3])
		user.Gender = record[4]
		user.Age, e = strconv.Atoi(record[5])
		if user.Age > 18 {
			user.Adult = true
		}

		Users = append(Users, user)
	}
}

func ParseFloat(vl string) float64 {

	s, err := strconv.ParseFloat(vl, 64)
	if err != nil {
		return 0
	}
	return s
}
