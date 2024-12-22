package service

import "log"
func CreateSchedule() Schedule {
    return Schedule{
        ID:          "1",
        StartTime:   "2024-12-22 09:00",
        EndTime:     "2024-12-22 10:00",
        Description: "Meeting with team",
    }
}

type PlandayService struct {}

func NewPlandayService() *PlandayService {
    return &PlandayService{}
}

func (s *PlandayService) FetchSchedules() ([]Schedule, error) {
    // Example: Fetch schedules from Planday API
    log.Println("Fetching schedules from Planday")
    return []Schedule{}, nil
}
