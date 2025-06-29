package api

import (
	"backend/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var issues = []entity.Issue{{Summary: "Убраться", Description: "Убраться в квартире", ID: 1, Status: "To Do"}, {Summary: "Помыть", Description: "Посуду или полы?", ID: 2, Status: "Done"}}

func Init(r *gin.Engine) {
	r.GET("/tasks", getIssue)
	r.POST("/tasks", postIssue)
	r.POST("/task/:ID", updateIssue)
	r.DELETE("/tasks/:ID", deleteIssue)
}

func getIssue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resp": issues,
	})
}

func postIssue(c *gin.Context) {
	var newIssue entity.Issue
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

func updateIssue(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var updatedIssue entity.Issue
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

func deleteIssue(c *gin.Context) {
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
