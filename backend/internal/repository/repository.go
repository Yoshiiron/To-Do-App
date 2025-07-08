package repository

import "backend/internal/domain"

type IssueRepository interface {
	Create(issue *domain.Issue) error
	Update(issue *domain.Issue, id int) error
	FindByID(id int) (*domain.Issue, int, error)
	ReturnAllIssues() ([]domain.Issue, error)
	Delete(id int) error
}
