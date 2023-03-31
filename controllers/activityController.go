package controllers

import (
	"net/http"
	"todo-app/database"
	"todo-app/helpers"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func CreateActivity(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Activity := models.Activity{}
	if contentType == appJSON {
		c.ShouldBindJSON(&Activity)
	} else {
		c.ShouldBind(&Activity)
	}

	err := db.Debug().Create(&Activity).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})

}

func GetAllActivity(c *gin.Context) {
	db := database.GetDB()
	Activities := []models.Activity{}

	err := db.Find(&Activities).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activities,
	})

}

func GetOneActivity(c *gin.Context) {
	db := database.GetDB()
	Activity := models.Activity{}
	ActivityId := c.Param("activityId")

	err := db.Where("activity_id = ?", ActivityId).First(&Activity).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})

}

func UpdateActivity(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Activity := models.Activity{}
	ActivityId := c.Param("activityId")
	ActivityUpdate := models.Activity{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Activity)
	} else {
		c.ShouldBind(&Activity)
	}

	err := db.Model(&Activity).Where("activity_id = ?", ActivityId).Updates(&Activity).Error
	_ = err

	db.Where("activity_id = ?", ActivityId).First(&ActivityUpdate)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    ActivityUpdate,
	})
}

func DeleteActivity(c *gin.Context) {
	db := database.GetDB()
	Activity := models.Activity{}
	ActivityId := c.Param("activityId")

	err := db.Where("activity_id = ?", ActivityId).Delete(&Activity).Error
	_ = err

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    gin.H{},
	})

}
