package routes

import (
	"github.com/davidgodeness/tododemo/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	g := r.Group("/api/v1")

	g.POST("/tasks", controller.CreateTask)
	g.GET("/tasks", controller.GetTasks)
	g.PATCH("/tasks/:id", controller.UpdateTask)
	g.DELETE("/tasks/:id", controller.DeleteTask)
	g.GET("/tasks/:id", controller.GetTask)
}
