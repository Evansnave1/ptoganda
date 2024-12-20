package services

import "log"

type PlandayService struct {}

func NewPlandayService() *PlandayService {
    return &PlandayService{}
}

func (s *PlandayService) FetchSchedules() ([]Schedule, error) {
    // Example: Fetch schedules from Planday API
    log.Println("Fetching schedules from Planday")
    return []Schedule{}, nil
}
