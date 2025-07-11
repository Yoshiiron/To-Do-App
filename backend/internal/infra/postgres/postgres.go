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
	db, err := gorm.Open(pg.Open("host=127.0.0.1 user=arb password=arb dbname=todo sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	tableExist := db.Migrator().HasTable("issues")
	if !tableExist {
		fmt.Println("Table not found, creating new 'Issues' table")
		CreateIssueTable(db)
	}

	return &issueRepository{db: db}
}

func CreateIssueTable(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&domain.Issue{})
	if err != nil {
		panic("Error creating table")
	}
	return nil
}

// Create implements repository.IssueRepository.
func (i *issueRepository) Create(issue *domain.Issue) error {
	i.db.Create(&issue)
	return nil
}

// Delete implements repository.IssueRepository.
func (i *issueRepository) Delete(id int) error {
	tx := i.db.Delete(&domain.Issue{}, id)
	return tx.Error
}

// FindByID implements repository.IssueRepository.
func (i *issueRepository) FindByID(id int) (*domain.Issue, int, error) {
	issue := &domain.Issue{}
	tx := i.db.Find(&issue, id)
	if issue.IssueID != 0 {
		return issue, issue.IssueID, tx.Error
	}
	return nil, 0, fmt.Errorf("no issue was found")
}

// ReturnAllIssues implements repository.IssueRepository.
func (i *issueRepository) ReturnAllIssues() ([]domain.Issue, error) {
	issues := []domain.Issue{}
	if len(issues) == 0 {
		return nil, fmt.Errorf("looks like db is empty")
	}
	return issues, nil
}

// Update implements repository.IssueRepository.
func (i *issueRepository) Update(issue *domain.Issue, id int) error {
	tx := i.db.Where("issue_id = ?", id).UpdateColumns(issue)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
