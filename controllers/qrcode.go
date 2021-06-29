package controllers

import (
	"fmt"
	"learning/db"

	"github.com/gin-gonic/gin"
	qrcode "github.com/yeqown/go-qrcode"
)

func CreatQrc(c *gin.Context) {
	url := c.Request.URL.Query()

	qrc, err := qrcode.New(db.Converter(url))
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err := qrc.Save("db/web-site.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
