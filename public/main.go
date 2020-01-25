package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	firstName string `json:"first_name"`
	surname string `json:"surname"`
	lastName string `json:"last_name"`
	password string `json:"password"`
} 

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		
	})
}
