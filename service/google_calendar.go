package service

import (
	"log"
	"google.golang.org/api/calendar/v3"
)

type Schedule struct {
	ID          string
	StartTime   string
	EndTime     string
	Description string
}

type GoogleCalendarService struct {
	calendarClient *calendar.Service
}

func NewGoogleCalendarService(client *calendar.Service) *GoogleCalendarService {
	return &GoogleCalendarService{calendarClient: client}
}

func (s *GoogleCalendarService) SyncWithPlanday(schedules []Schedule) error {
	log.Println("Syncing schedules with Google Calendar")
	return nil
}
