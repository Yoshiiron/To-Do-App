package api

import (
	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var issues = []models.Issue{{Summary: "Убраться", Description: "Убраться в квартире", ID: 1, Status: "To Do"}, {Summary: "Помыть", Description: "Посуду или полы?", ID: 2, Status: "Done"}}

func Init(r *gin.Engine) {
	r.GET("/tasks", getTasks)
	r.POST("/tasks", postTasks)
	r.POST("/task/:ID", updateTask)
	r.DELETE("/tasks/:ID", deleteTask)
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resp": issues,
	})
}

func postTasks(c *gin.Context) {
	var newIssue models.Issue
	if err := c.ShouldBindJSON(&newIssue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, issue := range issues {
		newIssue.ID = issue.ID + 1
	}

	newIssue.Status = "To Do"
	issues = append(issues, newIssue)

	c.JSON(http.StatusOK, gin.H{})
}

func updateTask(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var updatedIssue models.Issue
	if err := c.ShouldBindJSON(&updatedIssue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	for index, issue := range issues {
		if issue.ID == id {
			updatedIssue.ID = id
			issues[index] = updatedIssue
			c.JSON(http.StatusOK, gin.H{
				"result": updatedIssue,
			})
			break
		}
	}
}

func deleteTask(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for index, issue := range issues {
		if issue.ID == id {
			issues = append(issues[:index], issues[index+1:]...)
			c.JSON(http.StatusOK, gin.H{})
			break
		}
	}
}
