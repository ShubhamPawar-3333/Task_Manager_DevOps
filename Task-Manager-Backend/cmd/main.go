package main

import (
	"Task-Manager-Backend/internal/db"
	"Task-Manager-Backend/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize SQLite
	if err := db.InitSQLite(); err != nil {
		log.Fatal(err)
	}
	defer db.CloseSQLite()

	// Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Rate limiting middleware (if you have it)
	r.Use(handlers.RateLimitMiddleware())

	// Routes
	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}
