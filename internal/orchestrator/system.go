package orchestrator

import (
	"github.com/alexandravram45/spam-complaint/internal/complaint"
	"github.com/alexandravram45/spam-complaint/internal/detector"
	"github.com/alexandravram45/spam-complaint/internal/ingestor"
	"github.com/alexandravram45/spam-complaint/internal/logger"
)

type System struct {
	detector  *detector.SpamDetector
	complaint *complaint.Service
	ingestor  *ingestor.Ingestor
	log       *loggerWrapper
}

type loggerWrapper struct {
	Println func(v ...any)
}

func NewSystem(det *detector.SpamDetector) *System {
	log := logger.New()

	repo := complaint.NewRepository("data/complaints.json")
	service := complaint.NewService(repo)

	return &System{
		detector:  det,
		complaint: service,
		ingestor:  ingestor.NewIngestor(),
		log:       &loggerWrapper{Println: log.Println},
	}
}

func (s *System) Run() {
	s.log.Println("System started")

	for _, msg := range s.ingestor.Fetch() {
		s.log.Println("Processing message:", msg)

		if s.detector.IsSpam(msg) {
			s.complaint.Create(msg)
		} else {
			s.log.Println("Message is not spam")
		}
	}
}

func (s *System) ComplaintService() *complaint.Service {
	return s.complaint
}

func (s *System) Detector() *detector.SpamDetector {
	return s.detector
}
