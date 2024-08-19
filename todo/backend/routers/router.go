package routers

import (
	"github.com/Matama091/go_exercise/todo/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}
