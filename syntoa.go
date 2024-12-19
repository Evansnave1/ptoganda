package main

import (
    "os"
    "time"

    "github.com/arran4/golang-ical"
)

func generateICSFile(events []Event) {
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
