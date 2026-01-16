package main

import (
	"backend/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()

	router.Run(":8080")
}
