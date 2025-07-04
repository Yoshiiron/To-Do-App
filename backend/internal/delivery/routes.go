package delivery

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *IssueHandler) {
	r.GET("/tasks", handler.ReturnAllIssues)
	r.POST("/tasks", handler.Create)
	r.POST("/task/:ID", handler.Update)
	r.DELETE("/tasks/:ID", handler.Delete)
}

// func (h *IssueHandler) getIssue(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"resp": issues,
// 	})
// }

// func postIssue(c *gin.Context) {
// 	var newIssue domain.Issue
// 	if err := c.ShouldBindJSON(&newIssue); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	for _, issue := range issues {
// 		newIssue.ID = issue.ID + 1
// 	}

// 	newIssue.Status = "To Do"
// 	issues = append(issues, newIssue)

// 	c.JSON(http.StatusOK, gin.H{})
// }

// func updateIssue(c *gin.Context) {
// 	idParam := c.Param("ID")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	var updatedIssue domain.Issue
// 	if err := c.ShouldBindJSON(&updatedIssue); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 	}
// 	for index, issue := range issues {
// 		if issue.ID == id {
// 			updatedIssue.ID = id
// 			issues[index] = updatedIssue
// 			c.JSON(http.StatusOK, gin.H{
// 				"result": updatedIssue,
// 			})
// 			break
// 		}
// 	}
// }

// func deleteIssue(c *gin.Context) {
// 	idParam := c.Param("ID")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	for index, issue := range issues {
// 		if issue.ID == id {
// 			issues = append(issues[:index], issues[index+1:]...)
// 			c.JSON(http.StatusOK, gin.H{})
// 			break
// 		}
// 	}
// }
