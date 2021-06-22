 package db


// import ("encoding/csv"
// 		"os"
// 		"fmt")

// func InitDB() {
// 	file, err := os.Open("db/users.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.Comma = ';'
	
// 	for {
// 		record, e := reader.Read()
// 		if e != nil {
// 			fmt.Println(e)
// 			break
// 		}
// 		fmt.Println(record)
// 	}
// }
