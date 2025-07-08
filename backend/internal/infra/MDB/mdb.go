package mdb

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
)

var issues = []domain.Issue{{Summary: "Убраться", Description: "Убраться в квартире", IssueID: 1, Status: "To Do"}, {Summary: "Помыть", Description: "Посуду или полы?", IssueID: 2, Status: "Done"}}

type issueRepository struct {
	db []domain.Issue
}

func NewIssueRepository() repository.IssueRepository {
	return &issueRepository{db: issues}
}

// Create implements repository.IssueRepository.
func (i *issueRepository) Create(newIssue *domain.Issue) error {
	for _, issue := range i.db {
		newIssue.IssueID = issue.IssueID + 1
	}

	i.db = append(i.db, *newIssue)

	_, _, err := i.FindByID(newIssue.IssueID)
	return err
}

// Update implements repository.IssueRepository.
func (i *issueRepository) Update(issue *domain.Issue, IssueID int) error {

	_, index, err := i.FindByID(IssueID)
	if err != nil {
		return fmt.Errorf("")
	}

	issue.IssueID = IssueID
	i.db[index] = *issue

	fmt.Println(issue)

	return nil
}

// Delete implements repository.IssueRepository.
func (i *issueRepository) Delete(IssueID int) error {
	for index, issue := range i.db {
		if issue.IssueID == IssueID {
			i.db = append(i.db[:index], i.db[index+1:]...)
			break
		}
	}
	return nil
}

// FindByID implements repository.IssueRepository.
func (i *issueRepository) FindByID(IssueID int) (*domain.Issue, int, error) {
	for index, issue := range i.db {
		if issue.IssueID == IssueID {
			return &issue, index, nil
		}
	}
	return nil, 0, fmt.Errorf("error while trying to find issue by id")
}

// ReturnAllIssues implements repository.IssueRepository.
func (i *issueRepository) ReturnAllIssues() ([]domain.Issue, error) {
	return i.db, nil
}
