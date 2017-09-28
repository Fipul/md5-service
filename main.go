package main

import (
	"crypto/md5"
	"fmt"

	"github.com/gin-gonic/gin"
)

type md5Input struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func NewService() *gin.Engine {
	r := gin.Default()
	r.POST("/md5", func(c *gin.Context) {
		var input md5Input
		if err := c.BindJSON(&input); err != nil {
			return
		}
		if len(input.Text) > 100 {
			c.AbortWithStatus(400)
			return
		}
		sum := fmt.Sprintf("%x%d", md5.Sum([]byte(
			fmt.Sprintf("%d%s", input.ID, input.Text))), input.ID%2)
		c.Data(200, "text/plain", []byte(sum))
	})
	return r
}

func main() {
	NewService().Run(":8080")
}
