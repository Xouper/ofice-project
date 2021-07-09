package controllers

import (
	"bytes"
	"fmt"
	"learning/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	qrcode "github.com/yeqown/go-qrcode"
)

func CreatQrc(c *gin.Context) {
	url := c.Request.URL.Query()

	qrc, err := qrcode.New(db.Converter(url))
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	test := bytes.Buffer{}

	qrc.SaveTo(&test)
	r := bytes.NewReader(test.Bytes())
	extraHeaders := map[string]string{}
	c.DataFromReader(200, r.Size(), "image/jpeg", r, extraHeaders)
}

func AddUserImage(c *gin.Context) {
	// ID := c.Param("id")
	// IdInString, err := strconv.Atoi(ID)
	// if err != err {
	// 	return
	// }
	// user := db.GetUserById(&IdInString)
	// fmt.Printf("%v", user)

	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, file := range files {

		c.SaveUploadedFile(file, "db/images/"+file.Filename)
	}

	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files.", len(files)))

	//user := db.GetUserById()
}

func ShowUserImage(c *gin.Context) {
	ID := c.Param("id")
	IdInt, err := strconv.Atoi(ID)
		if err == nil {
			fmt.Println(IdInt)
	}
	for index := range db.Users {
		if index == int(IdInt) {
	
		}
	}
}