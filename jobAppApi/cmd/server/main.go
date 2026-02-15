package main

import (
	"jobAppApi/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := repository.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database: %v", err)
	}

	r := gin.Default()

	r.Run(":9080")
}
