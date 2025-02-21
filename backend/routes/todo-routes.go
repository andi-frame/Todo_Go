package routes

import (
	"fmt"
	"net/http"

	"github.com/andi-frame/Todo_Go/database"
	"github.com/andi-frame/Todo_Go/models"
	"github.com/gin-gonic/gin"
)

func createTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo successfully created!", "data": todo})
}

func getAllTodo(c *gin.Context) {
	var todos []models.Todo

	if err := database.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todos retrieved successfully", "data": todos})
}

func getTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Todo with ID %s not found", id), "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo retrieved successfully", "data": todo})
}

func updateTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Todo with ID %s not found", id), "details": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Todo with ID %s updated successfully", id), "data": todo})
}

func deleteTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Todo with ID %s not found", id), "details": err.Error()})
		return
	}

	if err := database.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Todo with ID %s deleted successfully", id)})
}

func TodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todo")
	{
		todoGroup.POST("/", createTodo)
		todoGroup.GET("/", getAllTodo)
		todoGroup.GET("/:id", getTodoById)
		todoGroup.PUT("/:id", updateTodoById)
		todoGroup.DELETE("/:id", deleteTodoById)
	}
}
