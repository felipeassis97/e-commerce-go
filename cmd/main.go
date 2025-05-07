package main

import (
	"go-api/controller/routes"
	"go-api/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	conn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	routes.RegisterRoutes(server, conn)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong Edited num 2",
		})
	})
	err = server.Run(":8000")
	if err != nil {
		panic(err)
	}
}
