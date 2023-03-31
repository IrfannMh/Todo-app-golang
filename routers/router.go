package routers

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	activityRouter := router.Group("/activity-groups")
	{
		activityRouter.POST("/", controllers.CreateActivity)
		activityRouter.GET("/", controllers.GetAllActivity)
		activityRouter.GET("/:activityId", controllers.GetOneActivity)
		activityRouter.PUT("/:activityId", controllers.UpdateActivity)
		activityRouter.DELETE("/:activityId", controllers.DeleteActivity)

	}
	todoRouter := router.Group("/todo-items")
	{
		todoRouter.POST("/", controllers.CreateTodo)
		todoRouter.GET("/", controllers.GetAllTodo)
		todoRouter.GET("/:todoId", controllers.GetOneTodo)
		todoRouter.PUT("/:todoId", controllers.UpdateTodo)
		todoRouter.DELETE("/:todoId", controllers.DeleteTodo)

	}

	return router
}
