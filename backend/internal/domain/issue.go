package domain

type Issue struct {
	IssueID     int `gorm:"primaryKey;autoIncrement"`
	Summary     string
	Description string
	Status      string
}
