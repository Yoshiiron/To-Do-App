package usecases

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type IssueService struct {
	repo repository.IssueRepository
}

func NewIssueService(repo repository.IssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (s *IssueService) Create(issue *domain.Issue) error {
	return s.repo.Create(issue)
}

func (s *IssueService) Update(issue *domain.Issue, id int) error {
	return s.repo.Update(issue, id)
}

func (s *IssueService) FindByID(id int) (*domain.Issue, int, error) {
	return s.repo.FindByID(id)
}

func (s *IssueService) ReturnAllIssues() ([]domain.Issue, error) {
	return s.repo.ReturnAllIssues()
}

func (s *IssueService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
