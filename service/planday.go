package service

import "log"

type Schedule struct {
	ID          string
	StartTime   string
	EndTime     string
	Description string
}

type PlandayService struct{}

func NewPlandayService() *PlandayService {
	return &PlandayService{}
}

func (s *PlandayService) FetchSchedules() ([]Schedule, error) {
	log.Println("Fetching schedules from Planday")
	return []Schedule{
		{
			ID:          "1",
			StartTime:   "2024-12-22T09:00:00",
			EndTime:     "2024-12-22T10:00:00",
			Description: "Team Meeting",
		},
	}, nil
}

func (s *PlandayService) GenerateICSFile() string {
	events := []Event{
		{
			ID:    "1",
			Title: "Team Meeting",
			Start: "2024-12-22T09:00:00",
			End:   "2024-12-22T10:00:00",
		},
	}
	return GenerateICSFile(events)
}
