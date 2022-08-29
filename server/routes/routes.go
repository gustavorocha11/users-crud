package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavorocha11/users-crud/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.GET("/", controllers.ShowAllUsers)
			user.GET("/:id", controllers.ShowUsers)
			user.POST("/", controllers.CreateUsers)
			user.PUT("/:id", controllers.EditUsers)
			user.DELETE("/:id", controllers.DeleteUsers)
		}
	}

	return router
}
