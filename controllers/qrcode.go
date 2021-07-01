package controllers

import (
	"bytes"
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

	test := bytes.Buffer{}

	qrc.SaveTo(&test)
	r := bytes.NewReader(test.Bytes())
	extraHeaders := map[string]string{}
	c.DataFromReader(200, r.Size(), "image/jpeg", r, extraHeaders)
}
