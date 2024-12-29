package service

import (
	"fmt"
	"time"
	"github.com/arran4/golang-ical"
)

type Event struct {
	Title    string
	Date     string
	Location string
	ID       string
	Start    string // ISO 8601 string, e.g., "2024-12-27T09:00:00"
	End      string
}

func PrintEventDetails(event Event) {
	fmt.Println("Event: ", event.Title)
	fmt.Println("Date: ", event.Date)
	fmt.Println("Location: ", event.Location)
}

func GenerateICSFile(events []Event) string {
	calendar := ics.NewCalendar()
	layout := "2006-01-02T15:04:05"

	for _, event := range events {
		icsEvent := calendar.AddEvent(event.ID)

		startTime, err := time.Parse(layout, event.Start)
		if err != nil {
			fmt.Printf("Error parsing start time: %v\n", err)
			continue
		}

		endTime, err := time.Parse(layout, event.End)
		if err != nil {
			fmt.Printf("Error parsing end time: %v\n", err)
			continue
		}

		icsEvent.SetSummary(event.Title)
		icsEvent.SetStartAt(startTime)
		icsEvent.SetEndAt(endTime)
	}

	file, err := os.Create("schedule.ics")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return ""
	}
	defer file.Close()

	icsData := calendar.Serialize()
	_, err = file.Write([]byte(icsData))
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return ""
	}

	return icsData
}
