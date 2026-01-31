package complaint

import (
	"fmt"
	"time"
)

type Service struct {
	repo       *Repository
	complaints []Complaint
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(message string) {
	complaint := Complaint{
		ID:        len(s.complaints) + 1,
		Message:   message,
		Status:    "OPEN",
		CreatedAt: time.Now(),
	}

	s.complaints = append(s.complaints, complaint)
	_ = s.repo.Save(s.complaints)

	fmt.Println("🚨 Spam complaint saved:", complaint.ID)
}

func (s *Service) All() []Complaint {
	return s.complaints
}

func (s *Service) HandleMessage(message string, isSpam bool) {
	if isSpam {
		s.Create(message)
	}
}
