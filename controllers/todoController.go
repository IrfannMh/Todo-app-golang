package controllers

import (
	"net/http"
	"todo-app/database"
	"todo-app/helpers"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func CreateTodo(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Todo := models.Todo{}
	if contentType == appJSON {
		c.ShouldBindJSON(&Todo)
	} else {
		c.ShouldBind(&Todo)
	}

	err := db.Debug().Create(&Todo).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func GetAllTodo(c *gin.Context) {
	db := database.GetDB()
	Todos := []models.Todo{}

	err := db.Find(&Todos).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todos,
	})
}

func GetOneTodo(c *gin.Context) {
	db := database.GetDB()
	Todo := models.Todo{}
	TodoID := c.Param("todoId")

	err := db.Where("todo_id = ?", TodoID).First(&Todo).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func UpdateTodo(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Todo := models.Activity{}
	TodoID := c.Param("todoId")
	TodoUpdate := models.Todo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Todo)
	} else {
		c.ShouldBind(&Todo)
	}

	err := db.Model(&Todo).Where("todo_id = ?", TodoID).Updates(&Todo).Error
	_ = err

	db.Where("todo_id = ?", TodoID).First(&TodoUpdate)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    TodoUpdate,
	})
}

func DeleteTodo(c *gin.Context) {
	db := database.GetDB()
	Todo := models.Todo{}
	TodoID := c.Param("todoId")

	err := db.Where("todo_id = ?", TodoID).Delete(&Todo).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    gin.H{},
	})
}
