package service

import (
	"fmt"
	"os"

	"github.com/arran4/golang-ical"
)
func PrintEventDetails(event Event) {
    fmt.Println("Event: ", event.Title)
    fmt.Println("Date: ", event.Date)
    fmt.Println("Location: ", event.Location)
}
// GenerateICSFile creates an iCalendar file
func GenerateICSFile(events []Event) {
	calendar := ics.NewCalendar()
	for _, event := range events {
		icsEvent := calendar.AddEvent(event.ID)
		icsEvent.SetSummary(event.Title)
		icsEvent.SetStartAt(event.Start)
		icsEvent.SetEndAt(event.End)
	}

	file, _ := os.Create("schedule.ics")
	defer file.Close()
	file.Write([]byte(calendar.Serialize()))
}