package postgres

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type issueRepository struct {
	db *gorm.DB
}

func NewIssueRepository() repository.IssueRepository {
	db, err := gorm.Open(pg.Open("host=localhost user=arb password=arb dbname=todo sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}

	return &issueRepository{db: db}
}

// Create implements repository.IssueRepository.
func (i *issueRepository) Create(issue *domain.Issue) error {
	i.db.Create(&issue)
	return nil
}

// Delete implements repository.IssueRepository.
func (i *issueRepository) Delete(id int) error {
	tx := i.db.Delete(&domain.Issue{}, id)
	fmt.Println(tx)
	return tx.Error
}

// FindByID implements repository.IssueRepository.
func (i *issueRepository) FindByID(id int) (*domain.Issue, int, error) {
	panic("unimplemented")
}

// ReturnAllIssues implements repository.IssueRepository.
func (i *issueRepository) ReturnAllIssues() ([]domain.Issue, error) {
	panic("unimplemented")
}

// Update implements repository.IssueRepository.
func (i *issueRepository) Update(issue *domain.Issue, id int) error {
	panic("unimplemented")
}
