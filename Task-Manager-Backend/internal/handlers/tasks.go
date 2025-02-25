package handlers

import (
	"Task-Manager-Backend/internal/db"
	"Task-Manager-Backend/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5) // 1 req/sec, burst of 5

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetTasks(c *gin.Context) {
	tasks, err := db.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}
	task.ID = time.Now().Format("20060102150405")
	task.CreatedAt = time.Now()

	if err := db.InsertTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updateData struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}

	// Fetch the existing task to preserve all fields (including title)
	existingTask, err := db.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update only the completed field, keeping title and other fields
	existingTask.Completed = updateData.Completed

	if err := db.UpdateTask(id, existingTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated task with all fields
	updatedTask, err := db.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := db.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
