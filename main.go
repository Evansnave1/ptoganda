package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"       // Import Gin framework
	"ptoganda/service"              // Replace with your actual module name
)

type App struct {
	GoogleCalendarService *service.GoogleCalendarService
	PlandayService        *service.PlandayService
}

func main() {
	googleService := service.NewGoogleCalendarService(nil)
	plandayService := service.NewPlandayService()

	app := &App{
		GoogleCalendarService: googleService,
		PlandayService:        plandayService,
	}

	router := gin.Default()

	router.GET("/health", healthCheck)
	router.POST("/sync-google", app.syncGoogleCalendar)
	router.GET("/schedule.ics", app.serveICSFile)

	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (app *App) syncGoogleCalendar(c *gin.Context) {
	schedules, err := app.PlandayService.FetchSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules from Planday"})
		return
	}

	err = app.GoogleCalendarService.SyncWithPlanday(schedules)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync with Google Calendar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully synced schedules"})
}

func (app *App) serveICSFile(c *gin.Context) {
	icsData := app.PlandayService.GenerateICSFile()
	c.Header("Content-Type", "text/calendar")
	c.String(http.StatusOK, icsData)
}
