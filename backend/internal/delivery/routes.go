package delivery

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *IssueHandler) {
	r.GET("/tasks", handler.ReturnAllIssues)
	r.GET("/tasks/:ID", handler.FindByID)
	r.POST("/tasks", handler.Create)
	r.POST("/task/:ID", handler.Update)
	r.DELETE("/tasks/:ID", handler.Delete)
}
