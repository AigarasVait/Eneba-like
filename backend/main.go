package main

import (
	"backend/db"
	"backend/router"
)

func main() {
	db.Init()

	r := router.SetupRouter()
	r.Run(":8080")
}
