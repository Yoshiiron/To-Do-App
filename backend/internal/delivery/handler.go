package delivery

import (
	"backend/internal/domain"
	"backend/internal/usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IssueHandler struct {
	service *usecases.IssueService
}

func NewIssueHandler(s *usecases.IssueService) *IssueHandler {
	return &IssueHandler{service: s}
}

func (h *IssueHandler) Create(c *gin.Context) {
	var newIssue domain.Issue
	if err := c.ShouldBindJSON(&newIssue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newIssue.Status = "To Do"
	err := h.service.Create(&newIssue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *IssueHandler) FindByID(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "issue with this id not found",
		})
		return
	}
	issue, _, err := h.service.FindByID(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "issue with this id not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": issue,
	})
}

func (h *IssueHandler) Update(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedIssue domain.Issue

	if err := c.ShouldBindJSON(&updatedIssue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.service.Update(&updatedIssue, id)
}

func (h *IssueHandler) ReturnAllIssues(c *gin.Context) {
	db, err := h.service.ReturnAllIssues()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": db,
	})

}

func (h *IssueHandler) Delete(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
