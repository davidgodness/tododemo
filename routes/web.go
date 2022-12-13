package routes

import "github.com/gin-gonic/gin"

func RegisterWebRoutes(r *gin.Engine) {
	r.POST("/login")
	r.POST("/logout")
}
