package main

import (
    "context"
    "log"

    "google.golang.org/api/calendar/v3"
    "google.golang.org/api/option"
)

func syncToGoogleCalendar(token string, event calendar.Event) {
    ctx := context.Background()
    srv, err := calendar.NewService(ctx, option.WithTokenSource(token))
    if err != nil {
        log.Fatalf("Unable to retrieve Calendar client: %v", err)
    }

    calendarId := "primary"
    _, err = srv.Events.Insert(calendarId, &event).Do()
    if err != nil {
        log.Fatalf("Unable to create event: %v", err)
    }
    log.Println("Event created")
}
