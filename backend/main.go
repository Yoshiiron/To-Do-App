package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Issue struct {
	Summary     string
	Description string
	Status      string
	ID          int
}

var issues = []Issue{{Summary: "Убраться", Description: "Убраться в квартире", ID: 1, Status: "To Do"}, {Summary: "Помыть", Description: "Посуду или полы?", ID: 2, Status: "Done"}}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"resp": issues,
		})
	})
	r.POST("/tasks", func(c *gin.Context) {
		var newIssue Issue
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

	})
	r.POST("/task/:ID", func(c *gin.Context) {
		idParam := c.Param("ID")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var updatedIssue Issue
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

	})
	r.DELETE("/tasks/:ID", func(c *gin.Context) {
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

	})
	r.Run()
}
